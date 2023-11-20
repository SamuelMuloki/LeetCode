package solutions

import (
	"math"
	"sort"
)

func MinimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	l := 0
	r := k - 1
	ans := math.MaxInt32

	for r < len(nums) {
		ans = min(ans, nums[r]-nums[l])
		l++
		r++
	}

	return ans
}
