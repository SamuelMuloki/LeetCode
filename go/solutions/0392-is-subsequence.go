package solutions

func IsSubsequence(s string, t string) bool {
	j := len(s) - 1
	for i := len(t) - 1; i >= 0 && j >= 0; i-- {
		if t[i] == s[j] {
			j--
			continue
		}
	}

	return j < 0
}
