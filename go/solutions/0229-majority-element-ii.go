package solutions

import (
	"math"
	"sort"
)

func MajorityElement2(nums []int) []int {
	sort.Ints(nums)
	res, count, seen := []int{}, 1, math.MinInt
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			count++
		} else {
			count = 1
		}

		if seen != nums[i] && (len(nums)/3) < count {
			seen = nums[i]
			res = append(res, nums[i])
		}
	}

	return res
}
