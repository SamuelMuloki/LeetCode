package solutions

func FirstPalindrome(words []string) string {
	for i := range words {
		if isPalindrome(words[i]) {
			return words[i]
		}
	}

	return ""
}
