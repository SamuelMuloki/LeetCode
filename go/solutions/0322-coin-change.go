package solutions

import (
	"math"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for _, coin := range coins {
		for remain := 1; remain <= amount; remain++ {
			if remain >= coin {
				dp[remain] = utils.Min(dp[remain], dp[remain-coin]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
