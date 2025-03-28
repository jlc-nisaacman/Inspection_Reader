package parsers

import (
	"main/utils"
	"os"
	"reflect"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func ReadPDF(pdfPath string) (formType string, report map[string]string) {
	// utils.LogSafe("Opening PDF: %s\n", pdfPath)
	file, err := os.Open(pdfPath)
	if err != nil {
		utils.LogSafe(" |%s| Could not open PDF: %v", pdfPath, err)
		return "", nil
	}
	defer file.Close()

	config := api.LoadConfiguration()
	data, err := api.ExportForm(file, "", config)
	if err != nil {
		utils.LogSafe(" |%s| Could not export PDF fields: %v", pdfPath, err)
		return "", nil
	}

	formIdentifiers := map[string][]string{
		"Backflow Form":     {"bf type", "bf make", "bf model", "bf size", "bf sn"},
		"Backflow Form Alt": {"group11", "dropdown12", "text7", "dropdown13", "text14"},
		"Pump Test Form":    {"rpm1", "total flow 1", "rated rpm", "hp", "make", "model"},
		"Dry System Form":   {"air psi", "water psi", "time", "model", "make", "size"},
		"Inspection Form":   {"drain test line 1", "insp_freq", "phone", "residual 1", "insp_#"},
	}

	report = make(map[string]string)
	report["pdf_path"] = pdfPath
	extractedFields := make(map[string]bool)

	for _, form := range data.Forms {
		for _, text := range form.TextFields {
			name := strings.TrimSpace(strings.ToLower(text.Name))
			if text.Multiline {
				text.Value = strings.ReplaceAll(text.Value, "\r", " ")
			}
			report[name] = text.Value
			extractedFields[name] = true
		}
		for _, radio := range form.RadioButtonGroups {
			name := strings.TrimSpace(strings.ToLower(radio.Name))
			report[name] = radio.Value
			extractedFields[name] = true
		}
		for _, combo := range form.ComboBoxes {
			name := strings.TrimSpace(strings.ToLower(combo.Name))
			report[name] = combo.Value
			extractedFields[name] = true
		}
	}

	// Log each entry of the report for debuging purposes
	// for k, v := range report {
	// 	utils.LogSafe(" |%s|  %+v | %+v", pdfPath, k, v)
	// }

	// Determine form type based on unique fields
	formType = "Unknown"
	for name, identifiers := range formIdentifiers {
		matchCount := 0
		for _, field := range identifiers {
			normField := strings.TrimSpace(strings.ToLower(field))
			if extractedFields[normField] {
				matchCount++
			}
		}
		// If a majority of the unique fields are found, identify the form
		if matchCount > len(identifiers)/2 {
			formType = name
			break
		}
	}
	utils.LogSafe(" |%s| Detected Form Type: %s", pdfPath, formType)
	return formType, report
}

func MapForm[T any](report map[string]string) T {
	var form T
	formValue := reflect.ValueOf(&form).Elem()
	formType := formValue.Type()

	for i := range formValue.NumField() {
		field := formValue.Field(i)
		jsonTag := formType.Field(i).Tag.Get("json")

		// Handle alternative keys (e.g., json:"Remarks,alt=Text4")
		keys := []string{}
		for part := range strings.SplitSeq(jsonTag, ",") {
			if strings.HasPrefix(part, "alt=") {
				keys = append(keys, strings.TrimPrefix(part, "alt="))
			} else {
				keys = append(keys, part)
			}
		}

		// Find the first available key in the report map
		for _, key := range keys {
			if value, exists := report[key]; exists {
				if field.CanSet() {
					field.SetString(value)
					break
				}
			}
		}
	}
	return form
}
