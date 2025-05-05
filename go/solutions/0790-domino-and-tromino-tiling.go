package solutions

func NumTilings(n int) int {
	const MOD = 1e9 + 7
	if n < 3 {
		return n
	}
	dp := make([]int, n)
	dp[0], dp[1], dp[2] = 1, 2, 5
	for i := 3; i < n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % MOD
	}
	return dp[n-1]
}
