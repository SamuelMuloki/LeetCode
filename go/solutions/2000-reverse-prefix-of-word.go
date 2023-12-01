package solutions

func ReversePrefix(word string, ch byte) string {
	idx := 0
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			idx = i
			break
		}
	}

	runes := []rune(word)
	if idx != 0 {
		for i, j := 0, idx; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
	}

	return string(runes)
}
