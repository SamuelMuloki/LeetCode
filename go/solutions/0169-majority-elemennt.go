// https://leetcode.com/problems/majority-element/
package solutions

import (
	"sort"
)

func MajorityElement(nums []int) int {
	sort.Ints(nums)
	res, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		if (len(nums) / 2) < count {
			return nums[i]
		}
	}

	return res
}
