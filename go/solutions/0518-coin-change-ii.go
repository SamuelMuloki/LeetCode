package solutions

func Change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for remain := coin; remain <= amount; remain++ {
			dp[remain] += dp[remain-coin]
		}
	}

	return dp[amount]
}
