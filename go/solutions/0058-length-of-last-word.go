package solutions

func LengthOfLastWord(s string) int {
	i := len(s) - 1
	for !isValidWordChar(s[i]) {
		i--
	}

	count := 0
	for i >= 0 && isValidWordChar(s[i]) {
		count++
		i--
	}

	return count
}

func isValidWordChar(ch uint8) bool {
	return (ch >= 65 && ch <= 90) || (ch >= 97 && ch <= 122)
}
