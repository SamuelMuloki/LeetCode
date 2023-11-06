package solutions

import "strings"

func WordPattern(pattern string, s string) bool {
	pMap, sMap := make([]int, 256), make(map[string]int)
	sArr := strings.Split(s, " ")
	if len(pattern) != len(sArr) {
		return false
	}

	for i := range pattern {
		if pMap[pattern[i]] != sMap[sArr[i]] {
			return false
		}

		pMap[pattern[i]] = i + 1
		sMap[sArr[i]] = i + 1
	}

	return true
}
