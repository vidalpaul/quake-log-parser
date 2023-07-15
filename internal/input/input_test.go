package input

import (
	"io/ioutil"
	"os"
	"testing"
)

// Helper function to create a temporary test file with given content
func createTempFile(t *testing.T, content string) string {
	tmpfile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer tmpfile.Close()

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	return tmpfile.Name()
}

func TestGetAbsFilePath(t *testing.T) {
	// Test when a file path is provided as a command-line argument
	expectedPath := "/path/to/file.log"
	os.Args = []string{"app", expectedPath}

	absPath, err := GetAbsFilePath()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if absPath != expectedPath {
		t.Errorf("Expected path '%s', but got '%s'", expectedPath, absPath)
	}

}

func TestGetLogFileContent(t *testing.T) {
	// Create a temporary file with content
	content := "Log file content"
	tmpFilePath := createTempFile(t, content)
	defer os.Remove(tmpFilePath)

	// Test reading the content of the temporary file
	fileContent, err := GetLogFileContent(tmpFilePath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if fileContent != content {
		t.Errorf("Expected content '%s', but got '%s'", content, fileContent)
	}

	// Test when the file does not exist
	nonexistentFilePath := "nonexistent.log"
	_, err = GetLogFileContent(nonexistentFilePath)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
