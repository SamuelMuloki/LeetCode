package solutions

import (
	"math"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func MaxDotProduct(nums1 []int, nums2 []int) int {
	firstMax, secondMax := math.MinInt, math.MinInt
	firstMin, secondMin := math.MaxInt, math.MaxInt

	for _, num := range nums1 {
		firstMax = utils.Max(firstMax, num)
		firstMin = utils.Min(firstMin, num)
	}

	for _, num := range nums2 {
		secondMax = utils.Max(secondMax, num)
		secondMin = utils.Min(secondMin, num)
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
			dp[j] = utils.Max(take, utils.Max(prevDp[j], dp[j+1]))
		}

		prevDp = dp
	}

	return dp[0]
}
