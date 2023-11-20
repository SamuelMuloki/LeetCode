package solutions

import "math"

func MaximumProduct(nums []int) int {
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	min1, min2 := math.MaxInt32, math.MaxInt32

	for i := range nums {
		if nums[i] > max1 {
			max1, max2, max3 = nums[i], max1, max2
		} else if nums[i] > max2 {
			max2, max3 = nums[i], max2
		} else if nums[i] > max3 {
			max3 = nums[i]
		}

		if nums[i] < min1 {
			min1, min2 = nums[i], min1
		} else if nums[i] < min2 {
			min2 = nums[i]
		}
	}

	return max(max1*max2*max3, min1*min2*max1)
}
