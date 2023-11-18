package solutions

import (
	"math"
)

func MinimumSum2(nums []int) int {
	n := len(nums)
	leftMin := make([]int, n)
	leftMin[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i])
	}

	rightMin := make([]int, n)
	rightMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMin[i] = min(rightMin[i+1], nums[i])
	}

	res := math.MaxInt
	for i := 1; i < n-1; i++ {
		if nums[i] > leftMin[i-1] && nums[i] > rightMin[i+1] {
			res = min(res, leftMin[i-1]+nums[i]+rightMin[i+1])
		}
	}

	if res == math.MaxInt {
		return -1
	}

	return res
}
