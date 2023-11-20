package solutions

import (
	"math"
)

func MaxDotProduct(nums1 []int, nums2 []int) int {
	firstMax, secondMax := math.MinInt, math.MinInt
	firstMin, secondMin := math.MaxInt, math.MaxInt

	for _, num := range nums1 {
		firstMax = max(firstMax, num)
		firstMin = min(firstMin, num)
	}

	for _, num := range nums2 {
		secondMax = max(secondMax, num)
		secondMin = min(secondMin, num)
	}

	if firstMax < 0 && secondMin > 0 {
		return firstMax * secondMin
	}

	if firstMin > 0 && secondMax < 0 {
		return firstMin * secondMax
	}

	m := len(nums2) + 1
	var dp []int
	prevDp := make([]int, m)

	for i := len(nums1) - 1; i >= 0; i-- {
		dp = make([]int, m)
		for j := len(nums2) - 1; j >= 0; j-- {
			take := (nums1[i] * nums2[j]) + prevDp[j+1]
			dp[j] = max(take, max(prevDp[j], dp[j+1]))
		}

		prevDp = dp
	}

	return dp[0]
}
