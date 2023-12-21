package solutions

import "math"

func MaxProductDifference(nums []int) int {
	max1, max2 := math.MinInt, math.MinInt
	min1, min2 := math.MaxInt, math.MaxInt
	for _, num := range nums {
		if num > max1 {
			max1, max2 = num, max1
		} else if num > max2 {
			max2 = num
		}

		if num < min1 {
			min1, min2 = num, min1
		} else if num < min2 {
			min2 = num
		}
	}

	return (max1 * max2) - (min1 * min2)
}
