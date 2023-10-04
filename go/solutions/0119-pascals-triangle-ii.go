package solutions

func GetRow(rowIndex int) []int {
	dp := make([][]int, rowIndex+1)
	for i := range dp {
		dp[i] = make([]int, rowIndex+1)
	}

	dp[0][0] = 1
	for i := 1; i <= rowIndex; i++ {
		dp[i][0] = 1
		for j := 1; j < len(dp[i-1]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
		}
	}

	return dp[rowIndex]
}
