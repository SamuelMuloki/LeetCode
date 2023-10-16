package solutions

import (
	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func MinPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[m-1])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == m && j == n-1 {
			return 0
		}

		if i == m || j == n {
			return 1e9
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		dp[i][j] = utils.Min(dfs(i+1, j), dfs(i, j+1)) + grid[i][j]
		return dp[i][j]
	}

	return dfs(0, 0)
}
