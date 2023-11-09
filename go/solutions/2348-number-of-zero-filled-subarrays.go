package solutions

func ZeroFilledSubarray(nums []int) int64 {
	ans, currStreak := int64(0), int64(0)
	for i := range nums {
		if nums[i] == 0 {
			currStreak++
		} else {
			currStreak = 0
		}

		ans += currStreak
	}

	return ans
}
