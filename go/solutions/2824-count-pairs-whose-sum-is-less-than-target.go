package solutions

import "sort"

func CountPairs(nums []int, target int) int {
	sort.Ints(nums)
	ans, left, right := 0, 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] < target {
			ans += right - left
			left++
		} else {
			right--
		}
	}

	return ans
}
