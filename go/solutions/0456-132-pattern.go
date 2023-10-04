package solutions

import "math"

func Find132pattern(nums []int) bool {
	st := []int{}
	third := math.MinInt32

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < third {
			return true
		}

		for len(st) > 0 && st[len(st)-1] < nums[i] {
			third = st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, nums[i])
	}

	return false
}
