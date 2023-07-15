package parser

// Quicksort algorithm for sorting the leaderboard
func quicksortLeaderboard(leaderboard []int, players []string, killCount map[string]int, low, high int) {
	if low < high {
		pivotIndex := partition(leaderboard, players, killCount, low, high)
		quicksortLeaderboard(leaderboard, players, killCount, low, pivotIndex-1)
		quicksortLeaderboard(leaderboard, players, killCount, pivotIndex+1, high)
	}
}

// Partition function for Quicksort
func partition(leaderboard []int, players []string, killCount map[string]int, low, high int) int {
	pivot := killCount[players[leaderboard[high]-1]]
	i := low - 1

	for j := low; j < high; j++ {
		if killCount[players[leaderboard[j]-1]] >= pivot {
			i++
			leaderboard[i], leaderboard[j] = leaderboard[j], leaderboard[i]
		}
	}

	leaderboard[i+1], leaderboard[high] = leaderboard[high], leaderboard[i+1]
	return i + 1
}
