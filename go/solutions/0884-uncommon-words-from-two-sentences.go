package solutions

import "strings"

func UncommonFromSentences(s1 string, s2 string) []string {
	sMap := make(map[string]int)
	s1Arr, s2Arr := strings.Split(s1, " "), strings.Split(s2, " ")
	for _, word := range s1Arr {
		sMap[word]++
	}

	for _, word := range s2Arr {
		sMap[word]++
	}

	ans := make([]string, 0)
	for word, count := range sMap {
		if count == 1 {
			ans = append(ans, word)
		}
	}

	return ans
}
