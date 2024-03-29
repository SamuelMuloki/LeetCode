package solutions

func LengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	maxLen, i, ok := 0, 0, false
	ss := []rune(s)

	for j := 0; j < len(ss); j++ {
		_, ok = m[ss[j]]
		if ok {
			i = max(i, m[ss[j]])
		}
		maxLen = max(maxLen, j-i+1)
		m[ss[j]] = j + 1
	}

	return maxLen
}
