package solutions

func IsAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}

	for i := range words {
		if words[i][0] != s[i] {
			return false
		}
	}

	return true
}
