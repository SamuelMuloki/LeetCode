package solutions

import "strings"

func IsAnagram(s string, t string) bool {
	sCount := make(map[string]int)
	sArr := strings.Split(s, "")

	if len(s) != len(t) {
		return false
	}

	for _, val := range sArr {
		sCount[val] = sCount[val] + 1
	}

	for key, value := range sCount {
		if strings.Count(t, key) != value {
			return false
		}
	}
	return true
}
