// Package parser provides functions to parse the Quake log file.

package parser

import (
	"fmt"
	"strings"
	"sync"

	"github.com/vidalpaul/quake-log-parser/internal/pkg/data"
)

// NewMatch creates a new match and adds it to the matchs map.
func NewMatch(matchs map[string]*data.MatchData, matchNumber int) *data.MatchData {
	newMatch := data.MatchData{
		TotalKills:  0,
		Players:     make([]string, 0),
		KillCount:   make(map[string]int),
		Leaderboard: make(map[int]string),
		KillMeans:   make(map[string]int),
	}

	fillKillMeans(&newMatch.KillMeans)

	matchName := fmt.Sprintf("game_%02d", matchNumber)
	matchs[matchName] = &newMatch

	return &newMatch
}

func fillKillMeans(means *map[string]int) {
	(*means)["MOD_UNKNOWN"] = 0
	(*means)["MOD_SHOTGUN"] = 0
	(*means)["MOD_GAUNTLET"] = 0
	(*means)["MOD_MACHINEGUN"] = 0
	(*means)["MOD_GRENADE"] = 0
	(*means)["MOD_GRENADE_SPLASH"] = 0
	(*means)["MOD_ROCKET"] = 0
	(*means)["MOD_ROCKET_SPLASH"] = 0
	(*means)["MOD_PLASMA"] = 0
	(*means)["MOD_PLASMA_SPLASH"] = 0
	(*means)["MOD_RAILGUN"] = 0
	(*means)["MOD_LIGHTNING"] = 0
	(*means)["MOD_BFG"] = 0
	(*means)["MOD_BFG_SPLASH"] = 0
	(*means)["MOD_WATER"] = 0
	(*means)["MOD_SLIME"] = 0
	(*means)["MOD_LAVA"] = 0
	(*means)["MOD_CRUSH"] = 0
	(*means)["MOD_TELEFRAG"] = 0
	(*means)["MOD_FALLING"] = 0
	(*means)["MOD_SUICIDE"] = 0
	(*means)["MOD_TARGET_LASER"] = 0
	(*means)["MOD_TRIGGER_HURT"] = 0
	(*means)["MOD_NAIL"] = 0
	(*means)["MOD_CHAINGUN"] = 0
	(*means)["MOD_PROXIMITY_MINE"] = 0
	(*means)["MOD_KAMIKAZE"] = 0
	(*means)["MOD_JUICED"] = 0
	(*means)["MOD_GRAPPL"] = 0
}

func NewLeaderboard(match *data.MatchData) {
	leaderboard := make([]string, len(match.Players))
	copy(leaderboard, match.Players)

	// Order by kills using insertion sort
	for i := 1; i < len(leaderboard); i++ {
		key := leaderboard[i]
		j := i - 1
		for j >= 0 && match.KillCount[leaderboard[j]] < match.KillCount[key] {
			leaderboard[j+1] = leaderboard[j]
			j--
		}
		leaderboard[j+1] = key
	}

	// Assign sorted leaderboard back to match.Leaderboard
	for i, player := range leaderboard {
		match.Leaderboard[i+1] = player
	}
}

func Parse(log string) map[string]*data.MatchData {
	var waitgroup sync.WaitGroup
	var matchs map[string]*data.MatchData = make(map[string]*data.MatchData, 0)
	matchNumber := 0

	// Log lines as array
	var lines []string = strings.Split(log, "\n")

	// Iterate log lines
	for lineNumber, line := range lines {
		var line string = strings.TrimSpace(line)
		var tokens []string = strings.Split(line, " ")

		// Find next match
		if len(tokens) > 2 {
			if tokens[1] == "InitGame:" {
				// New Match
				matchNumber++
				waitgroup.Add(1)
				newMatch := NewMatch(matchs, matchNumber)

				// Extract the data in parallel processe
				go ExtractMatchData(newMatch, lines, lineNumber+1, &waitgroup)
			}
		}
	}

	waitgroup.Wait()
	return matchs
}

func ExtractMatchData(match *data.MatchData, lines []string, lineNumber int, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()

	// Iterate log lines of specific match
	for lineNumber < len(lines) {
		var line string = strings.TrimSpace(lines[lineNumber])
		var tokens []string = strings.Split(line, " ")

		if len(tokens) > 1 {
			switch tokens[1] {
			// Kill log line
			case "Kill:":
				RegisterKill(match, tokens)

			// Player log line
			case "ClientUserinfoChanged:":
				RegisterPlayer(match, tokens)

			// Another match's started
			case "InitGame:":
				NewLeaderboard(match)
				return
			}
		}

		lineNumber++
	}

	// EOF
	NewLeaderboard(match)
}
