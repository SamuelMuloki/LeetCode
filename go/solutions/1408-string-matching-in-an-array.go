package solutions

import "strings"

func StringMatching(words []string) []string {
	sMap := make(map[string]bool)
	for i, word := range words {
		for j, sub := range words {
			if i == j {
				continue
			}

			if strings.Contains(word, sub) {
				sMap[sub] = true
			}
		}
	}

	ans := make([]string, 0)
	for word := range sMap {
		ans = append(ans, word)
	}

	return ans
}
