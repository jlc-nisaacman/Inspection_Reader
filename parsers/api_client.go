package parsers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/models"
	"main/utils"
	"net/http"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

// APIClient handles communication with the inspection API
type APIClient struct {
	BaseURL    string
	UUID       string
	HTTPClient *http.Client
}

// NewAPIClient creates a new API client with configuration from environment variables
func NewAPIClient() (*APIClient, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	baseURL := os.Getenv("API_URL")
	uuid := os.Getenv("READER_UUID")

	if baseURL == "" {
		return nil, fmt.Errorf("API_URL not set in environment")
	}
	if uuid == "" {
		return nil, fmt.Errorf("READER_UUID not set in environment")
	}

	return &APIClient{
		BaseURL: baseURL,
		UUID:    uuid,
		HTTPClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}, nil
}

// APIResponse represents the response from batch endpoints
type APIResponse struct {
	Inserted int `json:"inserted"`
	Updated  int `json:"updated"`
	Total    int `json:"total"`
}

// structToMap converts a struct to a map[string]interface{} using struct field names as keys
// This ensures the API receives field names that match database column names
// Field names are converted to lowercase (snake_case) to match PostgreSQL convention
func structToMap(s interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(s)
	t := v.Type()

	// Iterate through all struct fields
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name

		// Only include exported fields (those starting with uppercase)
		if t.Field(i).PkgPath == "" {
			// Convert PascalCase field name to snake_case for database
			// Since struct fields already use underscores, just convert to lowercase
			dbFieldName := strings.ToLower(fieldName)

			// Convert field value to interface{}
			fieldValue := field.Interface()

			// Handle date fields - convert invalid/empty strings to nil for PostgreSQL
			if strValue, ok := fieldValue.(string); ok && strings.Contains(strings.ToLower(fieldName), "date") {
				// Trim whitespace
				strValue = strings.TrimSpace(strValue)

				// Convert empty strings, "N/A", "None", or other invalid date strings to nil
				if strValue == "" ||
					strings.EqualFold(strValue, "N/A") ||
					strings.EqualFold(strValue, "None") ||
					strings.EqualFold(strValue, "null") {
					result[dbFieldName] = nil
				} else {
					result[dbFieldName] = strValue
				}
			} else {
				result[dbFieldName] = fieldValue
			}
		}
	}

	return result
}

// convertStructSliceToMaps converts a slice of structs to a slice of maps
func convertStructSliceToMaps(slice interface{}) []map[string]interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil
	}

	result := make([]map[string]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = structToMap(v.Index(i).Interface())
	}

	return result
}

// sendBatchRequest sends a batch request to the specified endpoint
func (c *APIClient) sendBatchRequest(endpoint string, data interface{}) (*APIResponse, error) {
	// Marshal data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	// Create request
	url := c.BaseURL + endpoint
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.UUID)

	// Send request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &apiResp, nil
}

// sendIndividualRequest sends a single record to the specified endpoint
func (c *APIClient) sendIndividualRequest(endpoint string, data map[string]interface{}) error {
	// Marshal data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	// Create request
	url := c.BaseURL + endpoint
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.UUID)

	// Send request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// SendInspectionsConcurrent sends inspection forms to the API using concurrent individual requests
func (c *APIClient) SendInspectionsConcurrent(forms []models.InspectionForm, progressCallback func()) error {
	if len(forms) == 0 {
		utils.LogSafe("No inspection forms to send")
		return nil
	}

	utils.LogSafe("Sending %d inspection forms concurrently...", len(forms))

	var (
		successCount int
		failedCount  int
		mu           sync.Mutex
		wg           sync.WaitGroup
	)

	// Worker pool with 20 concurrent workers
	workerPool := make(chan struct{}, 20)

	// Process each form concurrently
	for _, form := range forms {
		wg.Add(1)
		workerPool <- struct{}{} // Acquire worker slot

		go func(f models.InspectionForm) {
			defer wg.Done()
			defer func() { <-workerPool }() // Release worker slot

			// Convert struct to map
			data := structToMap(f)

			// Send individual request
			err := c.sendIndividualRequest("/inspections", data)

			mu.Lock()
			if err != nil {
				failedCount++
				utils.LogSafe("Failed to send %s: %v", f.PDF_Path, err)
			} else {
				successCount++
			}
			mu.Unlock()

			// Call progress callback if provided
			if progressCallback != nil {
				progressCallback()
			}
		}(form)
	}

	// Wait for all workers to complete
	wg.Wait()

	utils.LogSafe("Inspections complete: %d succeeded, %d failed, %d total",
		successCount, failedCount, len(forms))

	return nil
}

// SendInspectionsBatch sends inspection forms to the API using batch endpoint (deprecated - kept for compatibility)
func (c *APIClient) SendInspectionsBatch(forms []models.InspectionForm) error {
	if len(forms) == 0 {
		utils.LogSafe("No inspection forms to send")
		return nil
	}

	utils.LogSafe("Sending %d inspection forms to API...", len(forms))
	// Convert structs to maps using struct field names
	data := convertStructSliceToMaps(forms)
	resp, err := c.sendBatchRequest("/inspections/batch", data)
	if err != nil {
		return fmt.Errorf("failed to send inspections batch: %w", err)
	}

	utils.LogSafe("Inspections batch complete: %d inserted, %d updated, %d total",
		resp.Inserted, resp.Updated, resp.Total)
	return nil
}

// SendDrySystemsConcurrent sends dry system forms to the API using concurrent individual requests
func (c *APIClient) SendDrySystemsConcurrent(forms []models.DryForm, progressCallback func()) error {
	if len(forms) == 0 {
		utils.LogSafe("No dry system forms to send")
		return nil
	}

	utils.LogSafe("Sending %d dry system forms concurrently...", len(forms))

	var (
		successCount int
		failedCount  int
		mu           sync.Mutex
		wg           sync.WaitGroup
	)

	// Worker pool with 20 concurrent workers
	workerPool := make(chan struct{}, 20)

	// Process each form concurrently
	for _, form := range forms {
		wg.Add(1)
		workerPool <- struct{}{} // Acquire worker slot

		go func(f models.DryForm) {
			defer wg.Done()
			defer func() { <-workerPool }() // Release worker slot

			// Convert struct to map
			data := structToMap(f)

			// Send individual request
			err := c.sendIndividualRequest("/dry-systems", data)

			mu.Lock()
			if err != nil {
				failedCount++
				utils.LogSafe("Failed to send %s: %v", f.PDF_Path, err)
			} else {
				successCount++
			}
			mu.Unlock()

			// Call progress callback if provided
			if progressCallback != nil {
				progressCallback()
			}
		}(form)
	}

	// Wait for all workers to complete
	wg.Wait()

	utils.LogSafe("Dry systems complete: %d succeeded, %d failed, %d total",
		successCount, failedCount, len(forms))

	return nil
}

// SendDrySystemsBatch sends dry system forms to the API using batch endpoint (deprecated - kept for compatibility)
func (c *APIClient) SendDrySystemsBatch(forms []models.DryForm) error {
	if len(forms) == 0 {
		utils.LogSafe("No dry system forms to send")
		return nil
	}

	utils.LogSafe("Sending %d dry system forms to API...", len(forms))
	// Convert structs to maps using struct field names
	data := convertStructSliceToMaps(forms)
	resp, err := c.sendBatchRequest("/dry-systems/batch", data)
	if err != nil {
		return fmt.Errorf("failed to send dry systems batch: %w", err)
	}

	utils.LogSafe("Dry systems batch complete: %d inserted, %d updated, %d total",
		resp.Inserted, resp.Updated, resp.Total)
	return nil
}

// SendPumpSystemsConcurrent sends pump system forms to the API using concurrent individual requests
func (c *APIClient) SendPumpSystemsConcurrent(forms []models.PumpForm, progressCallback func()) error {
	if len(forms) == 0 {
		utils.LogSafe("No pump system forms to send")
		return nil
	}

	utils.LogSafe("Sending %d pump system forms concurrently...", len(forms))

	var (
		successCount int
		failedCount  int
		mu           sync.Mutex
		wg           sync.WaitGroup
	)

	// Worker pool with 20 concurrent workers
	workerPool := make(chan struct{}, 20)

	// Process each form concurrently
	for _, form := range forms {
		wg.Add(1)
		workerPool <- struct{}{} // Acquire worker slot

		go func(f models.PumpForm) {
			defer wg.Done()
			defer func() { <-workerPool }() // Release worker slot

			// Convert struct to map
			data := structToMap(f)

			// Send individual request
			err := c.sendIndividualRequest("/pump-systems", data)

			mu.Lock()
			if err != nil {
				failedCount++
				utils.LogSafe("Failed to send %s: %v", f.PDF_Path, err)
			} else {
				successCount++
			}
			mu.Unlock()

			// Call progress callback if provided
			if progressCallback != nil {
				progressCallback()
			}
		}(form)
	}

	// Wait for all workers to complete
	wg.Wait()

	utils.LogSafe("Pump systems complete: %d succeeded, %d failed, %d total",
		successCount, failedCount, len(forms))

	return nil
}

// SendPumpSystemsBatch sends pump system forms to the API using batch endpoint (deprecated - kept for compatibility)
func (c *APIClient) SendPumpSystemsBatch(forms []models.PumpForm) error {
	if len(forms) == 0 {
		utils.LogSafe("No pump system forms to send")
		return nil
	}

	utils.LogSafe("Sending %d pump system forms to API...", len(forms))
	// Convert structs to maps using struct field names
	data := convertStructSliceToMaps(forms)
	resp, err := c.sendBatchRequest("/pump-systems/batch", data)
	if err != nil {
		return fmt.Errorf("failed to send pump systems batch: %w", err)
	}

	utils.LogSafe("Pump systems batch complete: %d inserted, %d updated, %d total",
		resp.Inserted, resp.Updated, resp.Total)
	return nil
}

// SendBackflowConcurrent sends backflow forms to the API using concurrent individual requests
func (c *APIClient) SendBackflowConcurrent(forms []models.BackflowForm, progressCallback func()) error {
	if len(forms) == 0 {
		utils.LogSafe("No backflow forms to send")
		return nil
	}

	utils.LogSafe("Sending %d backflow forms concurrently...", len(forms))

	var (
		successCount int
		failedCount  int
		mu           sync.Mutex
		wg           sync.WaitGroup
	)

	// Worker pool with 20 concurrent workers
	workerPool := make(chan struct{}, 20)

	// Process each form concurrently
	for _, form := range forms {
		wg.Add(1)
		workerPool <- struct{}{} // Acquire worker slot

		go func(f models.BackflowForm) {
			defer wg.Done()
			defer func() { <-workerPool }() // Release worker slot

			// Convert struct to map
			data := structToMap(f)

			// Send individual request
			err := c.sendIndividualRequest("/backflow", data)

			mu.Lock()
			if err != nil {
				failedCount++
				utils.LogSafe("Failed to send %s: %v", f.PDF_Path, err)
			} else {
				successCount++
			}
			mu.Unlock()

			// Call progress callback if provided
			if progressCallback != nil {
				progressCallback()
			}
		}(form)
	}

	// Wait for all workers to complete
	wg.Wait()

	utils.LogSafe("Backflow complete: %d succeeded, %d failed, %d total",
		successCount, failedCount, len(forms))

	return nil
}

// SendBackflowBatch sends backflow forms to the API using batch endpoint (deprecated - kept for compatibility)
func (c *APIClient) SendBackflowBatch(forms []models.BackflowForm) error {
	if len(forms) == 0 {
		utils.LogSafe("No backflow forms to send")
		return nil
	}

	utils.LogSafe("Sending %d backflow forms to API...", len(forms))
	// Convert structs to maps using struct field names
	data := convertStructSliceToMaps(forms)
	resp, err := c.sendBatchRequest("/backflow/batch", data)
	if err != nil {
		return fmt.Errorf("failed to send backflow batch: %w", err)
	}

	utils.LogSafe("Backflow batch complete: %d inserted, %d updated, %d total",
		resp.Inserted, resp.Updated, resp.Total)
	return nil
}
