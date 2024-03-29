package solutions

func PaintWalls(cost []int, time []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	prevDp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prevDp[i] = 1e9
	}

	for i := n - 1; i >= 0; i-- {
		dp = make([]int, n+1)
		for remain := 1; remain <= n; remain++ {
			paint := cost[i] + prevDp[max(0, remain-1-time[i])]
			dontPaint := prevDp[remain]
			dp[remain] = min(paint, dontPaint)
		}
		prevDp = dp
	}

	return dp[n]
}
