package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func MaxProfit(prices []int) int {
	profit := 0

	i, j := 0, 1
	for j < len(prices) {
		if prices[i] < prices[j] {
			profit = utils.Max(profit, prices[j]-prices[i])
		} else {
			i = j
		}
		j++
	}

	return profit
}
