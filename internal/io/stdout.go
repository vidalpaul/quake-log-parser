package io

import (
	. "github.com/vidalpaul/quake-log-parser/internal/pkg/data"
)

func PrintReport(matchs map[string]*MatchData) {
	for index, match := range matchs {
		println("-------------------------- " + index + "Report --------------------------")
		println("TotalKills:", match.TotalKills)
		for player, kills := range match.KillCount {
			println(player, kills)
		}
	}
}
