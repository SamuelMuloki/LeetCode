package solutions

func CountGoodStrings(low int, high int, zero int, one int) int {
	const MOD = 1000000007
	dp := make([]int, high+1)

	for i := low; i <= high; i++ {
		dp[i] = 1
	}

	for value := high; value >= 0; value-- {
		if value+zero <= high {
			dp[value] = (dp[value] + dp[value+zero]) % MOD
		}
		if value+one <= high {
			dp[value] = (dp[value] + dp[value+one]) % MOD
		}
	}

	return dp[0]
}
