package parser_test

import (
	"testing"

	"github.com/vidalpaul/quake-log-parser/internal/parser"
	"github.com/vidalpaul/quake-log-parser/internal/pkg/data"
)

func TestNewLeaderboard(t *testing.T) {
	match := &data.MatchData{
		TotalKills:  0,
		Players:     []string{"Player1", "Player2", "Player3"},
		KillCount:   map[string]int{"Player1": 5, "Player2": 3, "Player3": 7},
		Leaderboard: make(map[int]string),
		KillMeans:   make(map[string]int),
	}

	parser.NewLeaderboard(match)

	expectedLeaderboard := map[int]string{
		1: "Player3",
		2: "Player1",
		3: "Player2",
	}

	for position, expectedPlayer := range expectedLeaderboard {
		if player, ok := match.Leaderboard[position]; !ok || player != expectedPlayer {
			t.Errorf("Expected player %s at position %d, got %s", expectedPlayer, position, player)
		}
	}
}
