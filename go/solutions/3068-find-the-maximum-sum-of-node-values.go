package solutions

func MaximumValueSum(nums []int, k int, edges [][]int) int64 {
	sum := int64(0)
	minExtra := int64(1000000)
	count := 0

	for _, val := range nums {
		if (val ^ k) > val {
			sum += int64(val ^ k)
			minExtra = min(minExtra, int64((val^k)-val))
			count++
		} else {
			sum += int64(val)
			minExtra = min(minExtra, int64(val-(val^k)))
		}
	}

	if count%2 == 0 {
		return sum
	} else {
		return sum - minExtra
	}
}
