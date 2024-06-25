package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	vtBaseURL  = "https://www.virustotal.com/vtapi/v2/"
	apiKeyFile = ".vtapikey" // File to store the API key locally
)

// CheckAPIKey validates the provided API key by making a simple request to VirusTotal.
func CheckAPIKey(apiKey string) bool {
	url := vtBaseURL + "file/report?apikey=" + apiKey + "&resource=dummy"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	if result["response_code"].(float64) == 0 {
		return true
	}
	return false
}

// GetAPIKey prompts the user for the VirusTotal API key and validates it.
func GetAPIKey() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your VirusTotal API key: ")
		apiKey, _ := reader.ReadString('\n')
		apiKey = apiKey[:len(apiKey)-1] // Remove newline character

		if CheckAPIKey(apiKey) {
			return apiKey
		} else {
			fmt.Println("Invalid API key, please try again.")
		}
	}
}

// UploadFile uploads the file to VirusTotal for scanning.
func UploadFile(apiKey, filePath string) (map[string]interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := bufio.NewWriter(body)
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	writer.Close()

	request, err := http.NewRequest("POST", vtBaseURL+"file/scan", body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("apikey", apiKey)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	json.Unmarshal(respBody, &result)
	return result, nil
}

func main() {

	apiKey := GetAPIKey()

	fmt.Print("Enter the path to the file you want to scan: ")
	reader := bufio.NewReader(os.Stdin)
	filePath, _ := reader.ReadString('\n')
	filePath = filePath[:len(filePath)-1] // Remove newline character

	result, err := UploadFile(apiKey, filePath)
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}

	fmt.Println("Scan Results:")
	fmt.Println(result)
}
