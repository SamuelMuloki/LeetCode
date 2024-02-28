package solutions

func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, maxMove+1)
		}
	}

	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(N, i, j int) int
	dfs = func(N, i, j int) int {
		if i == m || j == n || i < 0 || j < 0 {
			return 1
		}

		if N == 0 {
			return 0
		}

		if dp[i][j][N] >= 0 {
			return dp[i][j][N]
		}

		dp[i][j][N] = ((dfs(N-1, i-1, j)+dfs(N-1, i+1, j))%MOD + (dfs(N-1, i, j-1)+dfs(N-1, i, j+1))%MOD) % MOD

		return dp[i][j][N]
	}

	return dfs(maxMove, startRow, startColumn)
}
