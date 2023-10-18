package solutions

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r == m || c == n || obstacleGrid[r][c] == 1 {
			return 0
		}

		if r == m-1 && c == n-1 {
			if obstacleGrid[r][c] == 1 {
				return 0
			}
			return 1
		}

		if dp[r][c] != 0 {
			return dp[r][c]
		}

		dp[r][c] = dfs(r+1, c) + dfs(r, c+1)
		return dp[r][c]
	}

	return dfs(0, 0)
}
