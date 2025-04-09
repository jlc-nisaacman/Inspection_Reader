# Inspections Reader
Collects fillable PDFs from PDF_PATH. Identify which report the PDF is, then writes the data to the database.

# Environment Varables
```bash
DB_USER=
DB_PASS=
DB_HOST=
DB_NAME=
PDF_PATH=
```
# Logos
[icons8.com](https://icons8.com/license)

# Building Instructions
Compile logo
```bash
windres.exe icon.rc -O coff icon.syso
```

Compile program
```bash
go build -ldflags "-s -w" -o .\bin\inspections_reader.exe
```
