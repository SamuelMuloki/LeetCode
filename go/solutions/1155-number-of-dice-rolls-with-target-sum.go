package solutions

func NumRollsToTarget(n int, k int, target int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			for x := 1; x <= k; x++ {
				if x <= j {
					dp[i][j] = (dp[i][j] + dp[i-1][j-x]) % (1000000007)
				}
			}
		}
	}
	return dp[n][target]
}
