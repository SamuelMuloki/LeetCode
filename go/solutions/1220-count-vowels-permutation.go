package solutions

func CountVowelPermutation(n int) int {
	mod := 1000000007
	dp := make([][]int, 5)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[0][1], dp[1][1], dp[2][1], dp[3][1], dp[4][1] = 1, 1, 1, 1, 1
	for i := 2; i <= n; i++ {
		dp[0][i] = dp[1][i-1]
		dp[1][i] = (dp[0][i-1] + dp[2][i-1]) % mod
		dp[2][i] = (dp[0][i-1] + dp[1][i-1] + dp[3][i-1] + dp[4][i-1]) % mod
		dp[3][i] = (dp[2][i-1] + dp[4][i-1]) % mod
		dp[4][i] = dp[0][i-1]
	}

	ans := 0
	for i := 0; i < 5; i++ {
		ans = (ans + dp[i][n]) % mod
	}

	return ans
}
