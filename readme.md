# Inspections Reader
Collects fillable PDFs from PDF_PATH. Identify which report the PDF is, then writes the data to the database.

# Environment Varables
DB_USER=
DB_PASS=
DB_HOST=
DB_NAME=
PDF_PATH=

# Building Instructions

```bash
go build -ldflags "-s -w" -o .\bin\inspections.exe
```
