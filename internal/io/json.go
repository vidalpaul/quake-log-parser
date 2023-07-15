package io

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteJsonToFile(data interface{}, fileName string) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		println("Error marshalling to JSON:", err)
		return err
	}

	filePath := filepath.Join("./reports", fileName+".json")
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Error creating JSON file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		println("Error writing JSON to file:", err)
		return err
	}

	return nil
}
