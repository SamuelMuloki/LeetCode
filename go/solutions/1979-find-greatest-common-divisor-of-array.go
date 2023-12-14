package solutions

import "slices"

func FindGCD(nums []int) int {
	x, y := slices.Min(nums), slices.Max(nums)
	for y > 0 {
		x, y = y, x%y
	}

	return x
}
