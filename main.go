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
		log.Fatalf("Error initializing API client: %+v", err)
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

	// Create a progress bar to show processing status (PDF parsing + API upload)
	// We'll update it manually: fileCount for parsing, then 4 more for API batches
	fileCount := len(pdfFiles)
	totalSteps := fileCount + 4 // PDFs + 4 batch API calls
	bar := pb.StartNew(totalSteps)

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
				bar.Increment() // Update progress bar even for skipped files
				return          // Skip if PDF couldn't be read properly
			}

			// Process the PDF based on its form type and add to batch
			switch formType {
			case "Inspection Form":
				// Map the report data to an InspectionForm struct
				form := parsers.MapForm[models.InspectionForm](report)
				mu.Lock()
				inspectionForms = append(inspectionForms, form)
				mu.Unlock()
			case "Dry System Form":
				form := parsers.MapForm[models.DryForm](report)
				mu.Lock()
				dryForms = append(dryForms, form)
				mu.Unlock()
			case "Pump Test Form":
				form := parsers.MapForm[models.PumpForm](report)
				mu.Lock()
				pumpForms = append(pumpForms, form)
				mu.Unlock()
			case "Backflow Form":
				form := parsers.MapForm[models.BackflowForm](report)
				parsers.ProcessBackflowChoices(&form)
				mu.Lock()
				backflowForms = append(backflowForms, form)
				mu.Unlock()
			case "Backflow Form Alt":
				form := parsers.MapForm[models.BackflowForm](report)
				parsers.ProcessBackflowChoices(&form)
				mu.Lock()
				backflowForms = append(backflowForms, form)
				mu.Unlock()
			default:
				utils.LogSafe(" |%s| Skipping unknown form type.", pdf)
			}

			bar.Increment() // Update progress bar after processing each PDF
		}(pdf)
	}

	// Wait for all workers to complete
	wg.Wait()

	// Send all batches to the API
	utils.LogSafe("\n=== Sending batches to API ===")

	if err := apiClient.SendInspectionsBatch(inspectionForms); err != nil {
		log.Printf("Error sending inspections batch: %+v", err)
	}
	bar.Increment()

	if err := apiClient.SendDrySystemsBatch(dryForms); err != nil {
		log.Printf("Error sending dry systems batch: %+v", err)
	}
	bar.Increment()

	if err := apiClient.SendPumpSystemsBatch(pumpForms); err != nil {
		log.Printf("Error sending pump systems batch: %+v", err)
	}
	bar.Increment()

	if err := apiClient.SendBackflowBatch(backflowForms); err != nil {
		log.Printf("Error sending backflow batch: %+v", err)
	}
	bar.Increment()

	// Finish and close the progress bar
	bar.Finish()

	utils.LogSafe("=== Processing complete ===")
}
