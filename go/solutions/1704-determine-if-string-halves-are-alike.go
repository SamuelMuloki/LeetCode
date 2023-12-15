package solutions

func HalvesAreAlike(s string) bool {
	n := len(s)
	l, r := 0, 0
	for i := 0; i < n/2; i++ {
		if isVowel(rune(s[i])) {
			l++
		}

		if isVowel(rune(s[n-i-1])) {
			r++
		}
	}

	return l == r
}
