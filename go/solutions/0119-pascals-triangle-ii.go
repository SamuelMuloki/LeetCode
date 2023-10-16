package solutions

func GetRow(rowIndex int) []int {
	dp := make([]int, rowIndex+1)
	dp[0] = 1
	for i := 0; i < rowIndex; i++ {
		for j := i + 1; j > 0; j-- {
			dp[j] += dp[j-1]
		}
	}

	return dp
}
