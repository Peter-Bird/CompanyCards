package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// processCSV reads the CSV file, formats each record, and writes the output to individual text files.
func processCSV(inputFile, outputDir string) error {
	// Ensure the output directory exists
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Open the CSV file
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	// Read the header row to map field names to indices
	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read header: %v", err)
	}
	headerMap := make(map[string]int)
	for i, header := range headers {
		headerMap[header] = i
	}

	// Process each CSV record
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading record: %v", err)
		}

		// Retrieve the necessary fields using the header map
		symbol := record[headerMap["Symbol"]]
		security := record[headerMap["Security"]]
		gicsSector := record[headerMap["GICS Sector"]]
		gicsSubIndustry := record[headerMap["GICS Sub-Industry"]]
		headquartersLocation := record[headerMap["Headquarters Location"]]
		dateAdded := record[headerMap["Date added"]]
		cik := record[headerMap["CIK"]]
		founded := record[headerMap["Founded"]]

		// Format the output string with each key-value pair on a new line.
		outputString := fmt.Sprintf(
			"Symbol: %s\nSecurity: %s\nGICS Sector: %s\nGICS Sub-Industry: %s\nHeadquarters Location: \"%s\"\nDate added: %s\nCIK: %s\nFounded: %s\n",
			symbol, security, gicsSector, gicsSubIndustry, headquartersLocation, dateAdded, cik, founded,
		)

		// Define the output file name
		outputFile := filepath.Join(outputDir, fmt.Sprintf("%s.txt", symbol))

		// Write to the output file
		if err := os.WriteFile(outputFile, []byte(outputString), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %v", outputFile, err)
		}
	}

	fmt.Printf("Processed entries. Files saved in '%s'.\n", outputDir)
	return nil
}

func main() {
	inputFile := "sap500Info.csv"
	outputDir := "cards"

	if err := processCSV(inputFile, outputDir); err != nil {
		log.Fatal(err)
	}
}
