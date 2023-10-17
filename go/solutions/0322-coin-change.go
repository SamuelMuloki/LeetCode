package solutions

import (
	"math"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func CoinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make(map[int]int, n)

	var dfs func(remain int) int
	dfs = func(remain int) int {
		if count, ok := dp[remain]; ok {
			return count
		}

		if remain == 0 {
			return 0
		}

		if remain < 0 {
			return math.MaxInt32
		}

		minCount := math.MaxInt32
		for _, coin := range coins {
			minCount = utils.Min(minCount, dfs(remain-coin)+1)
		}

		dp[remain] = minCount
		return dp[remain]
	}

	res := dfs(amount)
	if res == math.MaxInt32 {
		return -1
	}

	return res
}
