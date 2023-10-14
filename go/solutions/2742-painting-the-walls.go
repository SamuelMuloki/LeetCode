package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func PaintWalls(cost []int, time []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	prevDp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prevDp[i] = 1e9
	}

	prevDp[0] = 0
	for i := n - 1; i >= 0; i-- {
		dp = make([]int, n+1)
		for remain := 1; remain <= n; remain++ {
			paint := cost[i] + prevDp[utils.Max(0, remain-1-time[i])]
			dontPaint := prevDp[remain]
			dp[remain] = utils.Min(paint, dontPaint)
		}
		prevDp = dp
	}

	return dp[n]
}
