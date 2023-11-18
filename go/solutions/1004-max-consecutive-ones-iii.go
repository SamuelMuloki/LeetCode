package solutions

func LongestOnes(nums []int, k int) int {
	res, zeroCount := 0, 0
	currMax, pos := 0, 0
	for i := 0; i < len(nums); i++ {
		zeroCount += (nums[i] & 1) ^ 1

		if zeroCount > k {
			currMax -= (nums[pos] | 1)
			zeroCount -= (nums[pos] & 1) ^ 1
			pos++
		}

		currMax += nums[i] | 1
		res = max(res, currMax)
	}

	return res
}
