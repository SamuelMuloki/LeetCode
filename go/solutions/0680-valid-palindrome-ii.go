package solutions

func ValidPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return validPalindromeUtil(s, i+1, j) || validPalindromeUtil(s, i, j-1)
		}

		i++
		j--
	}

	return true
}

func validPalindromeUtil(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
