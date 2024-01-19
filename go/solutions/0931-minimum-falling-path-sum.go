package solutions

import "math"

func MinFallingPathSum(matrix [][]int) int {
	if len(matrix) == 1 {
		return matrix[0][0]
	}

	n := len(matrix)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		copy(dp[i], matrix[i])
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < n; j++ {
			minPath := dp[i+1][j]
			if j > 0 {
				minPath = min(minPath, dp[i+1][j-1])
			}
			if j < n-1 {
				minPath = min(minPath, dp[i+1][j+1])
			}
			dp[i][j] += minPath
		}
	}

	ans := math.MaxInt32
	for _, num := range dp[0] {
		ans = min(ans, num)
	}

	return ans
}
