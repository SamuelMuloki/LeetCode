package solutions

func CheckRecord(n int) int {
	const MOD = 1000000007

	switch n {
	case 1:
		return 3
	case 2:
		return 8
	case 3:
		return 19
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 2
	dp[2] = 4
	dp[3] = 7

	for i := 4; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % MOD
	}

	withA := 0
	for i := 0; i < n; i++ {
		withA = (withA + dp[i]*dp[n-1-i]) % MOD
	}

	return (dp[n] + withA) % MOD
}
