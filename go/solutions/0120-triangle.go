package solutions

import (
	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func MinimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[n-1] = triangle[n-1]
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[i][j] = utils.Min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}

	return dp[0][0]
}
