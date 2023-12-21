package solutions

import "strings"

func PrintVertically(s string) []string {
	ans := []string{}
	splitS := strings.Split(s, " ")
	sLen := -1
	for i := range splitS {
		sLen = max(sLen, len(splitS[i]))
	}
	for i := 0; i < sLen; i++ {
		curr := []byte{}
		for j := 0; j < len(splitS); j++ {
			if i >= len(splitS[j]) {
				curr = append(curr, ' ')
			} else {
				curr = append(curr, splitS[j][i])
			}
		}
		ans = append(ans, strings.TrimRight(string(curr), " "))
	}

	return ans
}
