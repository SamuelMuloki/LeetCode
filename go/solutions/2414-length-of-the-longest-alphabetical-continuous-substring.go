package solutions

func LongestContinuousSubstring(s string) int {
	ans, currStreak := 1, 1
	runes := []rune(s)
	for i := 1; i < len(runes); i++ {
		if runes[i]-1 == runes[i-1] {
			currStreak++
		} else {
			currStreak = 1
		}

		ans = max(ans, currStreak)
	}

	return ans
}
