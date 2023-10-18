package solutions

func ToLowerCase(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += rune(32)
		}
	}

	return string(runes)
}
