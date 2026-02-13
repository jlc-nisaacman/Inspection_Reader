// $env:CGO_ENABLED=0; go build -ldflags="-s -w" -o .\bin\inspections_reader.exe
package main

import (
	"log"
	"main/models"
	"main/parsers"
	"main/utils"
	"os"
	"sync"

	"github.com/cheggaaa/pb/v3"
	"github.com/joho/godotenv"
)

// main is the entry point of the application.
// It loads environment variables, sets up logging, discovers PDF files,
// processes them in parallel using a worker pool, and sends their data to the API in batches.
func main() {
	// Set up logging to write to the inspections.log file
	utils.SetupLogging("inspections.log")

	// Load .env file for environment variables (API credentials, paths, etc.)
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read the PDF directory path from environment variables
	dir := os.Getenv("PDF_PATH")

	// Get all PDF files recursively from the specified directory
	pdfFiles, err := utils.GetPDFFiles(dir)
	if err != nil {
		log.Fatalf("Error getting PDFs: %+v", err)
	}

	// Initialize API client
	apiClient, err := parsers.NewAPIClient()
	if err != nil {
		// Log to file
		utils.LogSafe("ERROR: Failed to initialize API client: %+v", err)
		// Print to console (os.Stderr)
		os.Stderr.WriteString("\n========================================\n")
		os.Stderr.WriteString("ERROR: API is not available or misconfigured\n")
		os.Stderr.WriteString("Details: " + err.Error() + "\n")
		os.Stderr.WriteString("========================================\n\n")
		log.Fatalf("Exiting due to API initialization failure")
	}

	// Look up the user ID from the UUID via the API
	readerUserID, err := apiClient.GetUserIDFromUUID()
	if err != nil {
		// Log to file
		utils.LogSafe("ERROR: Failed to look up user ID from UUID: %v", err)
		// Print to console (os.Stderr)
		os.Stderr.WriteString("\n========================================\n")
		os.Stderr.WriteString("ERROR: API is not available or authentication failed\n")
		os.Stderr.WriteString("Details: " + err.Error() + "\n")
		os.Stderr.WriteString("Please check:\n")
		os.Stderr.WriteString("  - API_URL is correct in .env file\n")
		os.Stderr.WriteString("  - READER_UUID is valid\n")
		os.Stderr.WriteString("  - API server is running and accessible\n")
		os.Stderr.WriteString("========================================\n\n")
		log.Fatalf("Exiting due to API authentication failure")
	}

	// Create batch accumulators for each form type
	var (
		inspectionForms []models.InspectionForm
		dryForms        []models.DryForm
		pumpForms       []models.PumpForm
		backflowForms   []models.BackflowForm
		mu              sync.Mutex // Mutex to protect concurrent access to slices
	)

	// Set up a worker pool with a maximum of 10 concurrent workers
	var wg sync.WaitGroup
	workerPool := make(chan struct{}, 10)

	// Create first progress bar for PDF reading/walking
	fileCount := len(pdfFiles)
	pdfBar := pb.New(fileCount)
	pdfBar.SetTemplate(pb.Simple)
	pdfBar.Set("prefix", "Reading Reports: ")
	pdfBar.Start()

	// Process each PDF file in parallel
	for _, pdf := range pdfFiles {
		wg.Add(1)
		workerPool <- struct{}{} // Acquire a worker slot (blocks if pool is full)

		// Start a goroutine to process this PDF
		go func(pdf string) {
			defer wg.Done()
			defer func() { <-workerPool }() // Release worker slot when done

			// Read the PDF and identify its form type
			formType, report := parsers.ReadPDF(pdf)
			if report == nil {
				pdfBar.Increment() // Update progress bar even for skipped files
				return             // Skip if PDF couldn't be read properly
			}

			// Process the PDF based on its form type and add to batch
			switch formType {
			case "Inspection Form":
				// Map the report data to an InspectionForm struct
				form := parsers.MapForm[models.InspectionForm](report)
				form.PDF_Needed = false
				form.Created_By = readerUserID
				mu.Lock()
				inspectionForms = append(inspectionForms, form)
				mu.Unlock()
			case "Dry System Form":
				form := parsers.MapForm[models.DryForm](report)
				form.PDF_Needed = false
				form.Created_By = readerUserID
				mu.Lock()
				dryForms = append(dryForms, form)
				mu.Unlock()
			case "Pump Test Form":
				form := parsers.MapForm[models.PumpForm](report)
				form.PDF_Needed = false
				form.Created_By = readerUserID
				mu.Lock()
				pumpForms = append(pumpForms, form)
				mu.Unlock()
			case "Backflow Form":
				form := parsers.MapForm[models.BackflowForm](report)
				parsers.ProcessBackflowChoices(&form)
				form.PDF_Needed = false
				form.Created_By = readerUserID
				mu.Lock()
				backflowForms = append(backflowForms, form)
				mu.Unlock()
			case "Backflow Form Alt":
				form := parsers.MapForm[models.BackflowForm](report)
				parsers.ProcessBackflowChoices(&form)
				form.PDF_Needed = false
				form.Created_By = readerUserID
				mu.Lock()
				backflowForms = append(backflowForms, form)
				mu.Unlock()
			default:
				utils.LogSafe(" |%s| Skipping unknown form type.", pdf)
			}

			pdfBar.Increment() // Update progress bar after processing each PDF
		}(pdf)
	}

	// Wait for all workers to complete
	wg.Wait()

	// Finish the PDF reading progress bar
	pdfBar.Finish()

	// Create second progress bar for API uploads
	totalRecords := len(inspectionForms) + len(dryForms) + len(pumpForms) + len(backflowForms)
	uploadBar := pb.New(totalRecords)
	uploadBar.SetTemplate(pb.Simple)
	uploadBar.Set("prefix", "Uploading to API: ")

	// Send all records to the API concurrently
	utils.LogSafe("\n=== Sending records to API ===")
	uploadBar.Start()

	// Send inspections concurrently
	if err := apiClient.SendInspectionsConcurrent(inspectionForms, func() {
		uploadBar.Increment()
	}); err != nil {
		log.Printf("Error sending inspections: %+v", err)
	}

	// Send dry systems concurrently
	if err := apiClient.SendDrySystemsConcurrent(dryForms, func() {
		uploadBar.Increment()
	}); err != nil {
		log.Printf("Error sending dry systems: %+v", err)
	}

	// Send pump systems concurrently
	if err := apiClient.SendPumpSystemsConcurrent(pumpForms, func() {
		uploadBar.Increment()
	}); err != nil {
		log.Printf("Error sending pump systems: %+v", err)
	}

	// Send backflow concurrently
	if err := apiClient.SendBackflowConcurrent(backflowForms, func() {
		uploadBar.Increment()
	}); err != nil {
		log.Printf("Error sending backflow: %+v", err)
	}

	// Finish and close the upload progress bar
	uploadBar.Finish()

	utils.LogSafe("=== Processing complete ===")
}
