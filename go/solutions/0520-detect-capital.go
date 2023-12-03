package solutions

import "unicode"

func DetectCapitalUse(word string) bool {
	lu := -1
	for i := 1; i < len(word); i++ {
		if lu == -1 {
			if unicode.IsUpper(rune(word[i])) {
				lu = 1
			} else {
				lu = 0
			}

			if !unicode.IsUpper(rune(word[0])) && lu == 1 {
				return false
			}

			continue
		}

		if (lu == 1 && !unicode.IsUpper(rune(word[i]))) || (lu == 0 && !unicode.IsLower(rune(word[i]))) {
			return false
		}
	}

	return true
}
