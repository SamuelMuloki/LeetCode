package solutions

import "math"

func MaxAbsoluteSum(nums []int) int {
	maxSum, currMaxSum, currMinSum := math.MinInt, 0, 0
	for i := 0; i < len(nums); i++ {
		currMaxSum = max(nums[i], nums[i]+currMaxSum)
		currMinSum = min(nums[i], nums[i]+currMinSum)
		maxSum = max(maxSum, max(abs(currMaxSum), abs(currMinSum)))
	}

	return maxSum
}
