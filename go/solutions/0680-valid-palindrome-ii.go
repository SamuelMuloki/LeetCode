package solutions

func ValidPalindrome(s string) bool {
	valid, i, j := validPalindromeUtil(s, 0, len(s)-1)
	if valid {
		return true
	}

	if valid, _, _ := validPalindromeUtil(s, i+1, j); valid {
		return true
	}

	if valid, _, _ := validPalindromeUtil(s, i, j-1); valid {
		return true
	}

	return false
}

func validPalindromeUtil(s string, i, j int) (bool, int, int) {
	for i < j {
		if s[i] != s[j] {
			return false, i, j
		}

		i++
		j--
	}

	return true, i, j
}
