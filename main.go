package main

import (
	"bufio"
	"fmt"
	"go-secure/modfiles"
	"os"
)

func ScanFile(filePath string, apiKey string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	return nil
}

func main() {
	var name, apiKey string

	// Attempt to read from file
	nameFromFile, apiKeyFromFile, err := modfiles.ReadFromFile()
	if err != nil {

		fmt.Print("Enter Your Name: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			name = scanner.Text()
		}
		fmt.Print("Enter Your API Key: ")
		if scanner.Scan() {
			apiKey = scanner.Text()
		}

		// Write name and apiKey to file
		err := modfiles.WriteToFile(name, apiKey)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	} else {
		// File exists, use data from file
		name = nameFromFile
		apiKey = apiKeyFromFile

		fmt.Println("\t\t----- Welcome -----")
		modfiles.AsciiArtPrint()
		fmt.Println("User Name:", name, "API Key:", apiKey)
	}

	// Example: Scan a file using VirusTotal API (replace with actual file path)
	var filePath string
	fmt.Print("Enter your File Path:")
	fmt.Scan(&filePath)

	err = modfiles.ScanFile(filePath, apiKey)
	if err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

}
