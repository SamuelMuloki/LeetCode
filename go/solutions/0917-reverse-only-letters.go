package solutions

func ReverseOnlyLetters(s string) string {
	i, j := 0, len(s)-1
	runes := []rune(s)
	for i < j {
		if isLetter(runes[i]) && isLetter(runes[j]) {
			runes[i], runes[j] = runes[j], runes[i]
			i++
			j--
		} else if isLetter(runes[i]) {
			j--
		} else if isLetter(runes[j]) {
			i++
		} else {
			i++
			j--
		}
	}
	return string(runes)
}

func isLetter(ch rune) bool {
	return (ch >= 65 && ch <= 90) || (ch >= 97 && ch <= 122)
}
