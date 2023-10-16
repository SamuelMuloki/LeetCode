package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func MinimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == n {
			return 0
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		lowerLeft := triangle[i][j] + dfs(i+1, j)
		lowerRight := triangle[i][j] + dfs(i+1, j+1)

		dp[i][j] = utils.Min(lowerLeft, lowerRight)
		return dp[i][j]
	}

	return dfs(0, 0)
}
