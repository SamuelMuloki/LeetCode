package solutions

func ResultsArray(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	n := len(nums)
	result := make([]int, n-k+1)
	for i := range result {
		result[i] = -1
	}

	consecutiveCount := 1
	for i := 0; i < n-1; i++ {
		if nums[i]+1 == nums[i+1] {
			consecutiveCount++
		} else {
			consecutiveCount = 1
		}

		if consecutiveCount >= k {
			result[i-k+2] = nums[i+1]
		}
	}

	return result
}
