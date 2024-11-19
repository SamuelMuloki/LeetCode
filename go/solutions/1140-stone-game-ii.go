package solutions

func StoneGameII(piles []int) int {
	n := len(piles)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	suffixSum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + piles[i]
	}

	for i := 0; i <= n; i++ {
		dp[i][n] = suffixSum[i]
	}

	for i := n - 1; i >= 0; i-- {
		for maxTillNow := n - 1; maxTillNow >= 1; maxTillNow-- {
			for x := 1; x <= 2*maxTillNow && i+x <= n; x++ {
				dp[i][maxTillNow] = max(dp[i][maxTillNow], suffixSum[i]-dp[i+x][max(maxTillNow, x)])
			}
		}
	}

	return dp[0][1]
}
