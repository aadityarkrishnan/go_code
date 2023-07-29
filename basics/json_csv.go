package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	// Check if the correct number of arguments are provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: csv2json <input_csv_file> <output_json_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Open the CSV file
	csvFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Create a CSV reader
	csvReader := csv.NewReader(csvFile)

	// Read the header row to get the field names
	header, err := csvReader.Read()
	if err != nil {
		fmt.Println("Error reading CSV header:", err)
		return
	}

	// Create a slice to store the records
	var records []map[string]string

	// Read the CSV records and store them as maps with field names as keys
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV record:", err)
			return
		}

		// Create a map to store the record
		recordMap := make(map[string]string)

		// Map each field name to its corresponding value
		for i, value := range record {
			recordMap[header[i]] = value
		}

		records = append(records, recordMap)
	}

	// Create the output JSON file
	jsonFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// Create a JSON encoder and write the records to the JSON file
	jsonEncoder := json.NewEncoder(jsonFile)
	jsonEncoder.SetIndent("", "  ") // Optional: Add indentation for a more readable format
	err = jsonEncoder.Encode(records)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("CSV to JSON conversion completed successfully.")
}
