package solutions

func ReplaceDigits(s string) string {
	runes := []rune(s)
	for i := 1; i < len(runes); i += 2 {
		runes[i] = shift(runes[i-1], int(runes[i]-'0'))
	}

	return string(runes)
}

func shift(ch rune, digit int) rune {
	return rune(int(ch-'a') + digit + 'a')
}
