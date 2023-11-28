package solutions

func WaysToSplitArray(nums []int) int {
	n := len(nums)
	totalSum := 0
	for i := range nums {
		totalSum += nums[i]
	}

	currSum, ans := 0, 0
	for i := 0; i < n-1; i++ {
		currSum += nums[i]
		if totalSum-currSum <= currSum {
			ans++
		}
	}

	return ans
}
