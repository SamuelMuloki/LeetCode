package solutions

func MaxPower(s string) int {
	maxLen, currLen, prev := 0, 0, ' '
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == prev {
			currLen++
		} else {
			currLen = 1
			prev = rune(s[i])
		}

		maxLen = max(maxLen, currLen)
	}

	return maxLen
}
