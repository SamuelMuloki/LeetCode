package solutions

import "strings"

func FindWordsContaining(words []string, x byte) []int {
	ans := make([]int, 0)
	for i := range words {
		if strings.IndexByte(words[i], x) > -1 {
			ans = append(ans, i)
		}
	}

	return ans
}
