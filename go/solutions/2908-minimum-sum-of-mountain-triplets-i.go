package solutions

import (
	"math"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func MinimumSum(nums []int) int {
	n := len(nums)
	leftMin := make([]int, n)
	leftMin[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMin[i] = utils.Min(leftMin[i-1], nums[i])
	}

	rightMin := make([]int, n)
	rightMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMin[i] = utils.Min(rightMin[i+1], nums[i])
	}

	res := math.MaxInt
	for i := 1; i < n-1; i++ {
		if nums[i] > leftMin[i-1] && nums[i] > rightMin[i+1] {
			res = utils.Min(res, leftMin[i-1]+nums[i]+rightMin[i+1])
		}
	}

	if res == math.MaxInt {
		return -1
	}

	return res
}
