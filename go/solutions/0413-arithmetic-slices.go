package solutions

func NumberOfArithmeticSlices(nums []int) int {
	ans, currStreak, first := 0, 0, nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-first {
			currStreak++
		} else {
			currStreak = 0
		}

		first = nums[i-1]
		ans += currStreak
	}

	return ans
}
