package solutions

import "sort"

func CountFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)

	lower_bound := func(value int) int64 {
		left := 0
		right := len(nums) - 1
		result := int64(0)
		for left < right {
			sum := nums[left] + nums[right]
			if sum < value {
				result += int64(right - left)
				left++
			} else {
				right--
			}
		}

		return result
	}

	return lower_bound(upper+1) - lower_bound(lower)
}
