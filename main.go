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

func main() {

	utils.SetupLogging("inspections.log")

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read environment variables
	dir := os.Getenv("PDF_PATH")

	pdfFiles, err := utils.GetPDFFiles(dir)
	if err != nil {
		log.Fatalf("Error getting PDFs: %+v", err)
	}

	fileCount := len(pdfFiles)
	bar := pb.StartNew(fileCount)

	db := parsers.ConnectDatabase()
	defer db.Close()

	var wg sync.WaitGroup
	workerPool := make(chan struct{}, 10)

	for _, pdf := range pdfFiles {
		wg.Add(1)
		workerPool <- struct{}{} // Acquire a worker slot

		go func(pdf string) {
			defer wg.Done()
			defer func() { <-workerPool }() // Release worker slot

			bar.Increment()
			formType, report := parsers.ReadPDF(pdf)
			if report == nil {
				return
			}
			switch formType {
			case "Inspection Form":
				form := parsers.MapForm[models.InspectionForm](report)
				parsers.CreateInspectionTable(db)
				parsers.InsertInspectionTable(db, form)
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
	wg.Wait()
	bar.Finish()
}
