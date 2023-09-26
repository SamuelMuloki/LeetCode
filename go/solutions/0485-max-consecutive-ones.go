package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func FindMaxConsecutiveOnes(nums []int) int {
	res, currentMax := 0, 0
	for i := 0; i < len(nums); i++ {
		currentMax += nums[i] & 1
		res = utils.Max(currentMax, res)
		if nums[i]&1 == 0 {
			currentMax = 0
		}
	}

	return res
}
