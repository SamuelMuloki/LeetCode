package solutions

func DoesAliceWin(s string) bool {
	for i := 0; i < len(s); i++ {
		if isVowel(rune(s[i])) {
			return true
		}
	}

	return false
}
