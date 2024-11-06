package solutions

import (
	"math/bits"
	"sort"
)

func CanSortArray(nums []int) bool {
	sort.SliceStable(nums, func(i, j int) bool {
		countI, countJ := bits.OnesCount(uint(nums[i])), bits.OnesCount(uint(nums[j]))
		return countI == countJ && nums[i] < nums[j]
	})

	return sort.SliceIsSorted(nums, func(i, j int) bool { return nums[i] < nums[j] })
}
