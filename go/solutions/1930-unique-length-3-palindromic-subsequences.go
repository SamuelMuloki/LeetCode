package solutions

func CountPalindromicSubsequence(s string) int {
	letters := make(map[byte]bool)
	for i := range s {
		letters[s[i]] = true
	}

	ans := 0
	for letter := range letters {
		i, j := -1, 0
		for k := 0; k < len(s); k++ {
			if s[k] == letter {
				if i == -1 {
					i = k
				}

				j = k
			}
		}

		between := make(map[byte]bool)
		for k := i + 1; k < j; k++ {
			between[s[k]] = true
		}

		ans += len(between)
	}

	return ans
}
