package solutions

func FindLengthOfLCIS(nums []int) int {
	maxCount, count, prev := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > prev {
			count++
		} else {
			count = 1
		}

		prev = nums[i]
		maxCount = max(maxCount, count)
	}

	return maxCount
}
