package solutions

func MaxLengthBetweenEqualCharacters(s string) int {
	ans := -1
	for left := 0; left < len(s); left++ {
		for right := left + 1; right < len(s); right++ {
			if s[left] == s[right] {
				ans = max(ans, right-left-1)
			}
		}
	}

	return ans
}
