package output

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteJsonToFile(t *testing.T) {
	// Test data
	data := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	// Test file name
	fileName := "testfile"

	// Call the function
	err := WriteJsonToFile(data, fileName)
	if err != nil {
		t.Fatalf("Unexpected writeJsonToFile error: %v", err)
	}

	// Read the written file
	filePath := filepath.Join("./reports/", fileName+".json")
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read the written file: %v", err)
	}

	// Unmarshal the JSON content
	var resultData map[string]interface{}
	err = json.Unmarshal(fileContent, &resultData)
	if err != nil {
		t.Fatalf("Failed to unmarshal the JSON content: %v", err)
	}

	// Compare the data
	if len(resultData) != len(data) {
		t.Errorf("Unexpected number of elements in the result data. Got %d, want %d", len(resultData), len(data))
	}

	for key, value := range data {
		if resultValue, ok := resultData[key]; ok {
			if resultValue != value {
				t.Errorf("Unexpected value for key '%s'. Got %v, want %v", key, resultValue, value)
			}
		} else {
			t.Errorf("Key '%s' not found in the result data", key)
		}
	}

	// Clean up the written file
	err = os.Remove(filePath)
	if err != nil {
		t.Fatalf("Failed to clean up the written file: %v", err)
	}

	// Clean up the reports directory
	err = os.Remove("./reports/")
	if err != nil {
		t.Fatalf("Failed to clean up the reports directory: %v", err)
	}
}
