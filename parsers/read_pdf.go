package parsers

import (
	"main/utils"
	"os"
	"reflect"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// ReadPDF opens a PDF file, extracts form field data, and identifies the form type.
//
// Parameters:
//   - pdfPath: File path to the PDF to be processed
//
// Returns:
//   - formType: Identified type of form (e.g., "Inspection Form", "Backflow Form")
//   - report: Map of form field names to their values, or nil if extraction failed
func ReadPDF(pdfPath string) (formType string, report map[string]string) {
	// Open the PDF file
	file, err := os.Open(pdfPath)
	if err != nil {
		utils.LogSafe(" |%s| Could not open PDF: %v", pdfPath, err)
		return "", nil
	}
	defer file.Close()

	// Configure and export PDF form fields using pdfcpu
	config := api.LoadConfiguration()
	data, err := api.ExportForm(file, "", config)
	if err != nil {
		utils.LogSafe(" |%s| Could not export PDF fields: %v", pdfPath, err)
		return "", nil
	}

	// Map of form identifiers - keys are form types, values are field names that typically appear in that form type
	// Used to determine what kind of form we're processing based on presence of specific fields
	formIdentifiers := map[string][]string{
		"Backflow Form":     {"bf type", "bf make", "bf model", "bf size", "bf sn"},
		"Backflow Form Alt": {"group11", "dropdown12", "text7", "dropdown13", "text14"},
		"Pump Test Form":    {"rpm1", "total flow 1", "rated rpm", "hp", "make", "model"},
		"Dry System Form":   {"air psi", "water psi", "time", "model", "make", "size"},
		"Inspection Form":   {"drain test line 1", "insp_freq", "phone", "residual 1", "insp_#"},
	}

	// Initialize the report map and add the PDF path
	report = make(map[string]string)
	report["pdf_path"] = pdfPath
	extractedFields := make(map[string]bool) // Tracks which fields were found in the PDF

	// Extract all form fields from the PDF
	for _, form := range data.Forms {
		// Process text fields
		for _, text := range form.TextFields {
			name := strings.TrimSpace(strings.ToLower(text.Name))
			// Handle multiline text fields by replacing carriage returns with spaces
			if text.Multiline {
				text.Value = strings.ReplaceAll(text.Value, "\r", " ")
			}
			report[name] = text.Value
			extractedFields[name] = true
		}

		// Process radio button groups
		for _, radio := range form.RadioButtonGroups {
			name := strings.TrimSpace(strings.ToLower(radio.Name))
			report[name] = radio.Value
			extractedFields[name] = true
		}

		// Process combo boxes (dropdown selections)
		for _, combo := range form.ComboBoxes {
			name := strings.TrimSpace(strings.ToLower(combo.Name))
			report[name] = combo.Value
			extractedFields[name] = true
		}
	}

	// Determine form type based on the presence of identifying fields
	formType = "Unknown"
	for name, identifiers := range formIdentifiers {
		matchCount := 0
		for _, field := range identifiers {
			normField := strings.TrimSpace(strings.ToLower(field))
			if extractedFields[normField] {
				matchCount++
			}
		}
		// If a majority of the identifier fields are found, classify the form as this type
		if matchCount > len(identifiers)/2 {
			formType = name
			break
		}
	}

	utils.LogSafe(" |%s| Detected Form Type: %s", pdfPath, formType)
	return formType, report
}

// MapForm maps a generic report map to a specific form struct using reflection.
// It uses struct tags to match report keys with struct fields.
//
// Type Parameter:
//   - T: The target struct type to map the report data to
//
// Parameters:
//   - report: Map of field names to values extracted from a PDF
//
// Returns:
//   - An instance of type T populated with values from the report
func MapForm[T any](report map[string]string) T {
	var form T
	formValue := reflect.ValueOf(&form).Elem()
	formType := formValue.Type()

	// Iterate through each field in the target struct
	for i := 0; i < formValue.NumField(); i++ {
		field := formValue.Field(i)
		if !field.CanSet() {
			continue // Skip fields that can't be set (unexported)
		}

		// Get the json tag for this field
		jsonTag := formType.Field(i).Tag.Get("json")
		if jsonTag == "" {
			continue // Skip fields without json tags
		}

		// Handle alternative keys specified in the json tag
		keys := []string{}

		// Split by comma to separate the main tag from alternatives
		tagParts := strings.Split(jsonTag, ",")

		// First part is always the main field name
		if tagParts[0] != "" && tagParts[0] != "-" {
			keys = append(keys, strings.ToLower(tagParts[0]))
		}

		// Check remaining parts for alt= prefix
		for _, part := range tagParts[1:] {
			if strings.HasPrefix(part, "alt=") {
				altKey := strings.TrimPrefix(part, "alt=")
				keys = append(keys, strings.ToLower(altKey))
			}
		}

		// Try each possible key until we find a match in the report
		for _, key := range keys {
			if value, exists := report[key]; exists {
				field.SetString(value)
				break
			}
		}
	}
	return form
}
