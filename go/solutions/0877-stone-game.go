package solutions

func StoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < len(piles); i++ {
		dp[i][i] = piles[i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] > 0
}
