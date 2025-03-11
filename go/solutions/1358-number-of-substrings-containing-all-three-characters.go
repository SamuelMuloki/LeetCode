package solutions

func NumberOfSubstrings(s string) int {
	n := len(s)
	chSet := make(map[byte]int)
	ans := 0
	start, end := 0, 0
	for end < n {
		chSet[s[end]]++
		for start < n && len(chSet) == 3 {
			ans += n - end
			chSet[s[start]]--
			if chSet[s[start]] == 0 {
				delete(chSet, s[start])
			}
			start++
		}
		end++
	}

	return ans
}
