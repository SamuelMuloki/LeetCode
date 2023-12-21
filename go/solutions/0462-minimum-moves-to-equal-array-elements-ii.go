package solutions

import "sort"

func MinMoves2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ans, mid := 0, nums[n/2]
	for i := 0; i < n/2; i++ {
		ans += mid - nums[i]
		ans += nums[n-i-1] - mid
	}
	return ans
}
