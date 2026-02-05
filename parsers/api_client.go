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

// SendInspectionsBatch sends a batch of inspection forms to the API
func (c *APIClient) SendInspectionsBatch(forms []models.InspectionForm) error {
	if len(forms) == 0 {
		utils.LogSafe("No inspection forms to send")
		return nil
	}

	utils.LogSafe("Sending %d inspection forms to API...", len(forms))
	resp, err := c.sendBatchRequest("/inspections/batch", forms)
	if err != nil {
		return fmt.Errorf("failed to send inspections batch: %w", err)
	}

	utils.LogSafe("Inspections batch complete: %d inserted, %d updated, %d total",
		resp.Inserted, resp.Updated, resp.Total)
	return nil
}

// SendDrySystemsBatch sends a batch of dry system forms to the API
func (c *APIClient) SendDrySystemsBatch(forms []models.DryForm) error {
	if len(forms) == 0 {
		utils.LogSafe("No dry system forms to send")
		return nil
	}

	utils.LogSafe("Sending %d dry system forms to API...", len(forms))
	resp, err := c.sendBatchRequest("/dry-systems/batch", forms)
	if err != nil {
		return fmt.Errorf("failed to send dry systems batch: %w", err)
	}

	utils.LogSafe("Dry systems batch complete: %d inserted, %d updated, %d total",
		resp.Inserted, resp.Updated, resp.Total)
	return nil
}

// SendPumpSystemsBatch sends a batch of pump system forms to the API
func (c *APIClient) SendPumpSystemsBatch(forms []models.PumpForm) error {
	if len(forms) == 0 {
		utils.LogSafe("No pump system forms to send")
		return nil
	}

	utils.LogSafe("Sending %d pump system forms to API...", len(forms))
	resp, err := c.sendBatchRequest("/pump-systems/batch", forms)
	if err != nil {
		return fmt.Errorf("failed to send pump systems batch: %w", err)
	}

	utils.LogSafe("Pump systems batch complete: %d inserted, %d updated, %d total",
		resp.Inserted, resp.Updated, resp.Total)
	return nil
}

// SendBackflowBatch sends a batch of backflow forms to the API
func (c *APIClient) SendBackflowBatch(forms []models.BackflowForm) error {
	if len(forms) == 0 {
		utils.LogSafe("No backflow forms to send")
		return nil
	}

	utils.LogSafe("Sending %d backflow forms to API...", len(forms))
	resp, err := c.sendBatchRequest("/backflow/batch", forms)
	if err != nil {
		return fmt.Errorf("failed to send backflow batch: %w", err)
	}

	utils.LogSafe("Backflow batch complete: %d inserted, %d updated, %d total",
		resp.Inserted, resp.Updated, resp.Total)
	return nil
}
