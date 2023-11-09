package solutions

func NumSub(s string) int {
	ans, currStreak, MOD := 0, 0, int(1e9+7)
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			currStreak++
		} else {
			currStreak = 0
		}

		ans = (ans + currStreak) % MOD
	}

	return ans
}
