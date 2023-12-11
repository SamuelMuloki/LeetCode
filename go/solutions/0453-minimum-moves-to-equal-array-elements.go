package solutions

import "math"

func MinMoves(nums []int) int {
	minVal := math.MaxInt32
	for i := range nums {
		minVal = min(minVal, nums[i])
	}

	ans := 0
	for i := range nums {
		ans += nums[i] - minVal
	}

	return ans
}
