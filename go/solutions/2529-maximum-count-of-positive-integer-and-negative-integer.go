package solutions

import "sort"

func MaximumCount(nums []int) int {
	n := len(nums)
	neg := sort.Search(n, func(k int) bool { return nums[k] >= 0 })
	pos := sort.Search(n, func(k int) bool { return nums[k] > 0 })

	return max(neg, n-pos)
}
