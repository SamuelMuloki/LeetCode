package solutions

func CountSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count++
		for j := i + 1; j < len(s); j++ {
			str := string(s[i]) + string(s[i+1:j+1])

			if isPalindrome(str) {
				count++
			}
		}
	}

	return count
}
