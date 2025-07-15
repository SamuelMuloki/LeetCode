package solutions

func IsValid3(word string) bool {
	n := len(word)
	if n < 3 {
		return false
	}

	var isVowel = func(c rune) bool {
		return (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U')
	}

	var isAlphaNum = func(c rune) bool {
		return (c >= '0' && c <= '9' || c >= 'a' && c <= 'z' ||
			c >= 'A' && c <= 'Z')
	}

	var isConsonant = func(c rune) bool {
		return (c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z')
	}

	hasVowel := false
	hasConsonant := false
	for _, ch := range word {
		if !isAlphaNum(ch) {
			return false
		} else if isVowel(ch) {
			hasVowel = true
		} else if isConsonant(ch) {
			hasConsonant = true
		}
	}

	return hasVowel && hasConsonant
}
