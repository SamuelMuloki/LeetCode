package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func LongestSubarray(nums []int) int {
	res, zeroCount, maxCount, pos := 0, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		zeroCount += (nums[i] & 1) ^ 1

		if zeroCount > 1 {
			zeroCount -= (nums[pos] & 1) ^ 1
			maxCount -= (nums[pos] & 1)
			pos++
		}

		maxCount += (nums[i] & 1)
		res = utils.Max(res, maxCount)
	}

	if zeroCount == 0 {
		res--
	}

	return res
}
