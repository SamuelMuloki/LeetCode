package solutions

import "sort"

func BuyChoco(prices []int, money int) int {
	sort.Ints(prices)
	ans, sum := money, prices[0]+prices[1]
	if sum <= ans {
		ans -= sum
	}

	return ans
}
