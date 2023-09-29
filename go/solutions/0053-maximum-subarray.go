package solutions

import (
	"math"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func MaxSubArray(nums []int) int {
	maxSum, currSum := math.MinInt, 0
	for i := 0; i < len(nums); i++ {
		currSum = utils.Max(nums[i], nums[i]+currSum)
		maxSum = utils.Max(maxSum, currSum)
	}

	return maxSum
}
