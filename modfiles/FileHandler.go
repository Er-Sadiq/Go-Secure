package modfiles

import (
	"fmt"
	"os"
)

func WriteToFile(name, apiKey string) error {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile("GoSecureKey.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Prepare data to write
	data := fmt.Sprintf("Name: %s\nAPI Key: %s\n", name, apiKey)

	// Write data to file
	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func ReadFromFile() (string, string, error) {
	// Check if file exists
	_, err := os.Stat("GoSecureKey.txt")
	if os.IsNotExist(err) {
		return "", "", fmt.Errorf("File does not exist")
	}

	// Read file content
	content, err := os.ReadFile("GoSecureKey.txt")
	if err != nil {
		return "", "", err
	}

	// Parse content to extract name and API key
	var name, apiKey string
	fmt.Sscanf(string(content), "Name: %s\nAPI Key: %s\n", &name, &apiKey)
	return name, apiKey, nil
}
