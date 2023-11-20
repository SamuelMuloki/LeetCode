package solutions

func CountHomogenous(s string) int {
	ans, currStreak, MOD := 0, 0, int(1e9+7)
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] == s[i-1] {
			currStreak++
		} else {
			currStreak = 1
		}

		ans = (ans + currStreak) % MOD
	}

	return ans
}
