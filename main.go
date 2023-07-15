package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/vidalpaul/quake-log-parser/internal/input"
	"github.com/vidalpaul/quake-log-parser/internal/output"
	"github.com/vidalpaul/quake-log-parser/internal/parser"
)

func main() {
	start := time.Now()

	path, err := input.GetAbsFilePath()
	if err != nil {
		fmt.Println("Error getting file path:", err)
		return
	}

	println("Reading log file at", path, "...")
	content, err := input.GetLogFileContent(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	println("Parsing log data...")
	matchs := parser.Parse(content)

	println("Writing report to JSON file...")
	filename := filepath.Base(path)
	output.WriteReportToFile(matchs, filename)

	elapsed := time.Since(start)
	fmt.Printf("Parsing Quake log completed in %s\n", elapsed)
}
