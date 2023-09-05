package solutions

func ReverseVowels(s string) string {
	i, j := 0, len(s)-1
	runes := []rune(s)
	for i < j {
		if !isVowel(runes[i]) {
			i++
			continue
		}

		if !isVowel(runes[j]) {
			j--
			continue
		}

		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}

func isVowel(c rune) bool {
	return (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U')
}
