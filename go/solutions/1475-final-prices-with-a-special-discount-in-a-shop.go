package solutions

func FinalPrices(prices []int) []int {
	n := len(prices)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		curr := prices[i]
		for j := i + 1; j < n; j++ {
			if prices[j] <= curr {
				curr = curr - prices[j]
				break
			}
		}
		ans[i] = curr
	}

	return ans
}
