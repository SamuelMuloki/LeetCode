package solutions

import "unicode"

func CapitalizeTitle(title string) string {
	runes := []rune(title)
	k := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			k = i + 1
		}

		if unicode.IsUpper(runes[i]) {
			runes[i] += rune(32)
		}

		if i-k+1 > 2 && unicode.IsLower(runes[k]) {
			runes[k] -= rune(32)
		}
	}

	return string(runes)
}
