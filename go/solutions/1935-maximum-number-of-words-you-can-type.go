package solutions

func CanBeTypedWords(text string, brokenLetters string) int {
	broken := make(map[rune]bool)
	for _, ch := range brokenLetters {
		broken[ch] = true
	}

	res := 0
	found := false
	for _, ch := range text {
		if ch == ' ' {
			if !found {
				res++
			}
			found = false
		} else if broken[ch] {
			found = true
		}
	}

	if !found {
		res++
	}

	return res
}
