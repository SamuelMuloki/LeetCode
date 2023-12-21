package solutions

func AreOccurrencesEqual(s string) bool {
	set := [26]int{}
	cLen := 0
	for i := range s {
		set[s[i]-'a']++
		cLen = max(cLen, set[s[i]-'a'])
	}

	for i := 0; i < 26; i++ {
		if set[i] > 0 && set[i] != cLen {
			return false
		}
	}

	return true
}
