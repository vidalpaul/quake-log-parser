// File: register.go
// Package: parser
// This contains functions pertaining to registering kills and players.

package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/vidalpaul/quake-log-parser/internal/pkg/data"
)

func RegisterKill(match *data.MatchData, tokens []string) {
	// Adds total kills
	match.TotalKills++

	// Parse killer's name
	i := 5
	var killer string
	for tokens[i+1] != "killed" {
		killer += tokens[i] + " "
		i++
	}
	killer += tokens[i]
	i = i + 2

	// Parse victim's name
	var victim string
	for tokens[i+1] != "by" {
		victim += tokens[i] + " "
		i++
	}
	victim += tokens[i]
	i = i + 2

	// Parse kill mean
	var killMean string
	for i < len(tokens) {
		killMean += tokens[i]
		i++
	}

	if killer != "<world>" {
		// Register kill
		match.KillCount[killer]++
	} else {
		// Subtract <world>'s victim's kill count
		match.KillCount[victim]--
	}

	// Check if kill was by suicide or unknown mean
	if killer == victim {
		match.KillMeans["MOD_SUICIDE"]++
	} else if _, ok := match.KillMeans[killMean]; !ok {
		match.KillMeans["MOD_UNKNOWN"]++
	} else {
		match.KillMeans[killMean]++
	}
}

func RegisterPlayer(match *data.MatchData, tokens []string) {
	// Parse player name
	regex := regexp.MustCompile(`[^\\n](\w*|\w* )*`)
	player := regex.FindString(strings.Join(tokens[3:], " "))

	if len(player) > 1 {
		// Register new player
		if sliceContainsString(match.Players, player) {
			return
		} else {
			match.Players = append(match.Players, player)
			match.KillCount[player] = 0
		}
	} else {
		fmt.Println("No match found")
	}
}

func sliceContainsString(array []string, find string) bool {
	for _, aux := range array {
		if aux == find {
			return true
		}
	}
	return false
}
