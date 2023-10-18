package solutions

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid[0])
	dp := make([]int, n)

	dp[0] = 1
	for _, row := range obstacleGrid {
		for j := 0; j < n; j++ {
			if row[j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[n-1]
}
