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
// processes them in parallel using a worker pool, and stores their data in the database.
func main() {
	// Set up logging to write to the inspections.log file
	utils.SetupLogging("inspections.log")

	// Load .env file for environment variables (DB credentials, paths, etc.)
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

	// Create a progress bar to show processing status
	fileCount := len(pdfFiles)
	bar := pb.StartNew(fileCount)

	// Connect to the database
	db := parsers.ConnectDatabase()
	defer db.Close() // Ensure database connection is closed at the end

	// Set up a worker pool with a maximum of 10 concurrent workers
	var wg sync.WaitGroup
	workerPool := make(chan struct{}, 10)

	// Process each PDF file in parallel
	for _, pdf := range pdfFiles {
		wg.Add(1)
		workerPool <- struct{}{} // Acquire a worker slot (blocks if pool is full)

		// Start a goroutine to process this PDF
		go func(pdf string) {
			defer wg.Done()
			defer func() { <-workerPool }() // Release worker slot when done

			bar.Increment() // Update progress bar

			// Read the PDF and identify its form type
			formType, report := parsers.ReadPDF(pdf)
			if report == nil {
				return // Skip if PDF couldn't be read properly
			}

			// Process the PDF based on its form type
			switch formType {
			case "Inspection Form":
				// Map the report data to an InspectionForm struct
				form := parsers.MapForm[models.InspectionForm](report)
				parsers.CreateInspectionTable(db)       // Ensure table exists
				parsers.InsertInspectionTable(db, form) // Insert the data
			case "Dry System Form":
				form := parsers.MapForm[models.DryForm](report)
				parsers.CreateDryTable(db)
				parsers.InsertDryTable(db, form)
			case "Pump Test Form":
				form := parsers.MapForm[models.PumpForm](report)
				parsers.CreatePumpTable(db)
				parsers.InsertPumpTable(db, form)
			case "Backflow Form":
				form := parsers.MapForm[models.BackflowForm](report)
				parsers.CreateBackflowTable(db)
				parsers.InsertBackflowTable(db, form)
			case "Backflow Form Alt":
				form := parsers.MapForm[models.BackflowForm](report)
				parsers.CreateBackflowTable(db)
				parsers.InsertBackflowTable(db, form)
			default:
				utils.LogSafe(" |%s| Skipping unknown form type.", pdf)
			}
		}(pdf)
	}

	// Wait for all workers to complete
	wg.Wait()

	// Finish and close the progress bar
	bar.Finish()
}
