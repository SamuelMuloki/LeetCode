package solutions

func ScoreOfString(s string) int {
	ans := 0
	for i := 1; i < len(s); i++ {
		ans += abs(int(s[i]-'a') - int(s[i-1]-'a'))
	}

	return ans
}
