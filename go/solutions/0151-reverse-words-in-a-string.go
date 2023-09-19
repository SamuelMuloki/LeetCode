package solutions

import (
	"strings"
)

func ReverseWords(s string) string {
	words := make([]string, 0)
	k, end := len(s)-1, 0

	for k >= 0 && s[k] == ' ' {
		k--
	}

	end = k
	for k >= 0 {
		if s[k] != ' ' {
			k--
			continue
		}

		words = append(words, s[k+1:end+1])
		k--

		for k >= 0 && s[k] == ' ' {
			k--
		}

		end = k
	}

	if end >= 0 {
		words = append(words, s[:end+1])
	}

	return strings.Join(words, " ")
}
