package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func LengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	maxLen, i, ok := 0, 0, false
	ss := []rune(s)

	for j := 0; j < len(ss); j++ {
		_, ok = m[ss[j]]
		if ok {
			i = utils.Max(i, m[ss[j]])
		}
		maxLen = utils.Max(maxLen, j-i+1)
		m[ss[j]] = j + 1
	}

	return maxLen
}
