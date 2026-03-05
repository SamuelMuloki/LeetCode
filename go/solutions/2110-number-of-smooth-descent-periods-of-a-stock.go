package solutions

func GetDescentPeriods(prices []int) int64 {
	n := len(prices)
	res := int64(1)
	cur := 1
	for i := 1; i < n; i++ {
		if prices[i] == prices[i-1]-1 {
			cur++
		} else {
			cur = 1
		}
		res += int64(cur)
	}

	return res
}
