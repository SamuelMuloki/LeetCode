package solutions

func MaxProfit2(prices []int) int {
	ans := 0
	for i, j := 0, 1; j < len(prices); j++ {
		if prices[i] < prices[j] {
			ans += prices[j] - prices[i]
			i++
		} else {
			i = j
		}
	}

	return ans
}
