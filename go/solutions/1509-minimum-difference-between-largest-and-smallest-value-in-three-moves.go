package solutions

import (
	"math"
	"sort"
)

func MinDifference(nums []int) int {
	numsSize := len(nums)
	if numsSize <= 4 {
		return 0
	}

	sort.Ints(nums)
	minDiff := math.MaxInt
	for left, right := 0, numsSize-4; left < 4; left, right = left+1, right+1 {
		minDiff = min(minDiff, nums[right]-nums[left])
	}

	return minDiff
}
