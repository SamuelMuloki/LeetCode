package solutions

func NumberOfSubstrings(s string) int {
	n := len(s)
	chSet := [3]int{}
	start, ans := 0, 0
	for end := 0; end < n; end++ {
		chSet[s[end]-'a']++
		for chSet[0] > 0 && chSet[1] > 0 && chSet[2] > 0 {
			ans += (n - end)
			chSet[s[start]-'a']--
			start++
		}
	}

	return ans
}
