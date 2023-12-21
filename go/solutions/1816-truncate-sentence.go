package solutions

func TruncateSentence(s string, k int) string {
	for i := range s {
		if s[i] != ' ' {
			continue
		}

		if k == 1 {
			return string(s[:i])
		}

		k--
	}

	return s
}
