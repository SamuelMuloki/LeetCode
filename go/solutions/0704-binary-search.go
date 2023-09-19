// https://leetcode.com/problems/binary-search/
package solutions

import "math"

func Search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l != r {
		mid := int(math.Ceil(float64((l + r)) / 2))
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid
		}
	}

	if nums[l] == target {
		return l
	}

	return -1
}
