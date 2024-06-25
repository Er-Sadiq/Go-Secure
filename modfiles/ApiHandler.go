package modfiles

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

const (
	vtScanEndpoint = "https://www.virustotal.com/api/v3/files"
)

type VTScanResponse struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			LastAnalysisStats struct {
				Harmless   int `json:"harmless"`
				Malicious  int `json:"malicious"`
				Suspicious int `json:"suspicious"`
			} `json:"last_analysis_stats"`
			LastAnalysisResults map[string]interface{} `json:"last_analysis_results"`
		} `json:"attributes"`
	} `json:"data"`
}

func ScanFile(filePath string, apiKey string) error {
	client := resty.New()

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a multipart form request
	resp, err := client.R().
		SetHeader("x-apikey", apiKey).
		SetFileReader("file", filePath, file).
		Post(vtScanEndpoint)

	if err != nil {
		return fmt.Errorf("error making HTTP request: %v", err)
	}

	// Check the response status
	if resp.StatusCode() != 200 {
		return fmt.Errorf("unexpected response status: %v", resp.Status())
	}

	// Parse the response body
	var scanResponse VTScanResponse
	err = json.Unmarshal(resp.Body(), &scanResponse)
	if err != nil {
		return fmt.Errorf("error decoding response body: %v", err)
	}

	fmt.Println("Scan Results:")
	fmt.Printf("Harmless: %d\n", scanResponse.Data.Attributes.LastAnalysisStats.Harmless)
	fmt.Printf("Malicious: %d\n", scanResponse.Data.Attributes.LastAnalysisStats.Malicious)
	fmt.Printf("Suspicious: %d\n", scanResponse.Data.Attributes.LastAnalysisStats.Suspicious)

	return nil
}
