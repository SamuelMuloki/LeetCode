package solutions

import "math"

func MinMoves(nums []int) int {
	minVal, sum := math.MaxInt32, 0
	for i := range nums {
		sum += nums[i]
		minVal = min(minVal, nums[i])
	}

	return sum - len(nums)*minVal
}
