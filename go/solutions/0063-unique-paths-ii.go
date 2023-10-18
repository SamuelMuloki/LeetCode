package solutions

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for r := m - 1; r >= 0; r-- {
		for c := n - 1; c >= 0; c-- {
			if r == m-1 && c == n-1 {
				if obstacleGrid[r][c] == 1 {
					dp[r][c] = 0
				} else {
					dp[r][c] = 1
				}
			} else if r == m-1 || c == n-1 {
				if obstacleGrid[r][c] == 1 {
					dp[r][c] = 0
				} else if r == m-1 {
					dp[r][c] = dp[r][c+1]
				} else if c == n-1 {
					dp[r][c] = dp[r+1][c]
				}
			} else {
				if obstacleGrid[r][c] == 1 {
					dp[r][c] = 0
				} else {
					dp[r][c] = dp[r][c+1] + dp[r+1][c]
				}
			}
		}
	}

	return dp[0][0]
}
