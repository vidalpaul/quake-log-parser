package data

type MatchData struct {
	TotalKills  int            `json:"total_kills"`
	Players     []string       `json:"players"`
	KillCount   map[string]int `json:"kills"`
	Leaderboard map[int]string `json:"player_ranking"`
	KillMeans   map[string]int `json:"kills_by_means"`
}
