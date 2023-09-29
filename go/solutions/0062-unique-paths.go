package solutions

func UniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r == 1 || c == 1 {
			dp[r][c] = 1
			return dp[r][c]
		}

		if dp[r][c] == 0 {
			dp[r][c] = dfs(r-1, c) + dfs(r, c-1)
		}

		return dp[r][c]
	}

	return dfs(m, n)
}
