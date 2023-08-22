package solutions

func IsPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if !isValidChar(s[i]) {
			i++
			continue
		}

		if !isValidChar(s[j]) {
			j--
			continue
		}

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}

		i++
		j--
	}

	return true
}

func toLower(char uint8) uint8 {
	if 'A' <= char && char <= 'Z' {
		char += 'a' - 'A'
	}

	return char
}

func isValidChar(ch uint8) bool {
	return (ch >= 65 && ch <= 90) || (ch >= 48 && ch <= 57) || (ch >= 97 && ch <= 122)
}
