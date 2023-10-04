package solutions

func DecodeAtIndex(s string, k int) string {
	decodedLen := int64(0)
	runes := []rune(s)
	for _, ch := range runes {
		if isDigit(ch) {
			decodedLen *= int64(ch - '0')
		} else {
			decodedLen++
		}
	}

	for i := len(runes) - 1; i >= 0; i-- {
		if isDigit(runes[i]) {
			decodedLen /= int64(runes[i] - '0')
			k = k % int(decodedLen)
		} else {
			if k == 0 || int(decodedLen) == k {
				return string(runes[i])
			}
			decodedLen--
		}
	}

	return ""
}

func isDigit(d rune) bool {
	return (d >= 50 && d <= 57)
}
