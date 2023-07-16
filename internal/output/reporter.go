package output

import (
	"fmt"

	"github.com/vidalpaul/quake-log-parser/internal/pkg/data"
)

func WriteReportToFile(matchs map[string]*data.MatchData, filename string) {
	writeGroupedInformation(matchs, filename)
	fmt.Println("Report written to", filename+".json")
}

func writeGroupedInformation(matchs map[string]*data.MatchData, filename string) {
	WriteJsonToFile(matchs, filename)
}
