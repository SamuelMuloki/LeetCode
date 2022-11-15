package leetcode

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	nonAlphaNum := regexp.MustCompile("[^a-zA-Z0-9]+")
	str := strings.ToLower(nonAlphaNum.ReplaceAllString(s, ""))

	i, j := 0, len(str)-1
	for i <= j {
		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}

	return true
}
