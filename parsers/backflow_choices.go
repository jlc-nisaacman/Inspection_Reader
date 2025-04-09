package parsers

import (
	"main/models"
	"reflect"
	"strings"
)

// ProcessBackflowChoices modifies the BackflowForm struct to convert choice values
func ProcessBackflowChoices(form *models.BackflowForm) {
	choiceMappings := map[string]map[string]string{
		// Backflow Type mapping
		"Backflow_Type": {
			"choice1": "RPZ",
			"choice2": "DCVA",
			"choice3": "PVB",
			"choice4": "SRVB",
		},
		// Test Frequency mapping
		"Test_Type": {
			"choice5": "Test After Installation",
			"choice1": "Test After Repair",
			"choice2": "Annual",
			"choice3": "Semi Annual",
		},
		// Flow Condition Evaluated
		"PVB_SRVB_Check_Valve_Flow": {
			"choice2": "Flow",
			"choice1": "No-Flow",
		},
		// Downstream Shutoff Valve Status mapping
		"Downsteam_Shutoff_Valve_Status": {
			"choice2": "Closed Tight",
			"choice1": "Leaked",
			"choice3": "Not Tested",
		},
		// Service Type mapping
		"Protection_Type": {
			"choice4": "Service Line",
			"choice1": "Fire Service Line",
			"choice2": "Internal Domestic Plumbing System",
		},
		// Pass/Fail mapping
		"Result": {
			"choice3": "PASS",
			"choice1": "FAIL",
			"choice2": "OTHER",
		},
	}
	// Get the reflect.Value of the form
	v := reflect.ValueOf(form).Elem()

	// Iterate through the mappings
	for fieldName, mapping := range choiceMappings {
		// Find the field
		field := v.FieldByName(fieldName)
		if !field.IsValid() || field.Kind() != reflect.String {
			continue
		}

		// Get the current value
		currentValue := field.String()

		// Convert to lowercase for case-insensitive matching
		lowercaseValue := strings.ToLower(currentValue)

		// Check if the current value matches any choice key
		for choiceKey, choiceValue := range mapping {
			if lowercaseValue == strings.ToLower(choiceKey) {
				field.SetString(choiceValue)
				break
			}
		}
	}
}
