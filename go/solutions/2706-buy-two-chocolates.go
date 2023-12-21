package solutions

func BuyChoco(prices []int, money int) int {
	min1, min2 := 101, 101
	for i := range prices {
		if prices[i] < min1 {
			min1, min2 = prices[i], min1
		} else {
			min2 = min(min2, prices[i])
		}
	}

	if min1+min2 <= money {
		return money - (min1 + min2)
	}

	return money
}
