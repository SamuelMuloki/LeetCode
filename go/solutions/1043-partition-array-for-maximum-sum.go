package solutions

func MaxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		curMax := 0
		for j := 1; j <= min(k, i); j++ {
			curMax = max(curMax, arr[i-j])
			dp[i] = max(dp[i], dp[i-j]+curMax*j)
		}
	}
	return dp[n]
}
