package input

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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

func GetLogFileContent(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
