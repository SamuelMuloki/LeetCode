package solutions

import (
	"math"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func CoinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	for i := 1; i <= amount; i++ {
		dp[0][i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for remain := 1; remain <= amount; remain++ {
			if remain >= coins[i-1] {
				dp[i][remain] = utils.Min(dp[i-1][remain], dp[i][remain-coins[i-1]]+1)
			} else {
				dp[i][remain] = dp[i-1][remain]
			}
		}
	}

	if dp[n][amount] == math.MaxInt32 {
		return -1
	}

	return dp[n][amount]
}
