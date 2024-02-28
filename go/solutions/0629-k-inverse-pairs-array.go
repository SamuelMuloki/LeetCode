package solutions

func KInversePairs(n int, k int) int {
	dp := make([]int, k+1)
	dp[0] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[j] = (dp[j] + dp[j-1]) % MOD
		}
		for j := k; j >= i; j-- {
			dp[j] = (dp[j] - dp[j-i] + MOD) % MOD
		}
	}
	return dp[k]
}
