// Package input provides functions to get the log file path and its content.

package input

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// GetAbsFilePath returns the absolute path of the log file, or an error if it fails.
func GetAbsFilePath() (string, error) {
	var path string

	if len(os.Args) > 1 {
		path = strings.Join(os.Args[1:], " ")
	} else {
		// If no file path is provided, then use default file path
		path = "assets/logs/qgames.log"
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	return absPath, nil
}

// GetLogFileContent returns the content of the log file, or an error if it fails.
func GetLogFileContent(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
