package solutions

import "math"

func LongestSubarray3(nums []int) int {
	j := -1
	maxVal := math.MinInt32
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	res := 0
	for i, num := range nums {
		if num != maxVal {
			j = i
		}
		res = int(math.Max(float64(res), float64(i-j)))
	}
	return res
}
