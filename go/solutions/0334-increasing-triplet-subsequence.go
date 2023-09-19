package solutions

import "math"

func IncreasingTriplet(nums []int) bool {
	x, y := math.MaxInt, math.MaxInt
	for i := 0; i < len(nums); i++ {
		if nums[i] <= x {
			x = nums[i]
		} else if nums[i] < y {
			y = nums[i]
		} else if nums[i] > y {
			return true
		}
	}

	return false
}
