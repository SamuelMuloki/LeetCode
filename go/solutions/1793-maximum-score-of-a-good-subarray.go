package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func MaximumScore(nums []int, k int) int {
	n := len(nums)
	left, right := k, k
	ans, currMin := nums[k], nums[k]

	for left > 0 || right < n-1 {
		leftVal, rightVal := 0, 0
		if left > 0 {
			leftVal = nums[left-1]
		}

		if right < n-1 {
			rightVal = nums[right+1]
		}

		if leftVal < rightVal {
			right++
			currMin = utils.Min(currMin, nums[right])
		} else {
			left--
			currMin = utils.Min(currMin, nums[left])
		}

		ans = utils.Max(ans, currMin*(right-left+1))
	}

	return ans
}
