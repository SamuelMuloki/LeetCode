package solutions

import (
	"math"
)

func MaxSubArray(nums []int) int {
	maxSum, currSum := math.MinInt, 0
	for i := 0; i < len(nums); i++ {
		currSum = max(nums[i], nums[i]+currSum)
		maxSum = max(maxSum, currSum)
	}

	return maxSum
}
