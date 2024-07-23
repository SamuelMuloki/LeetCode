package solutions

import "math"

func MinFallingPathSum2(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		copy(dp[i], grid[i])
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < n; j++ {
			nextMin := math.MaxInt32
			for k := 0; k < n; k++ {
				if k != j {
					nextMin = min(nextMin, dp[i+1][k])
				}
			}
			dp[i][j] += nextMin
		}
	}

	ans := math.MaxInt32
	for _, num := range dp[0] {
		ans = min(ans, num)
	}

	return ans
}
