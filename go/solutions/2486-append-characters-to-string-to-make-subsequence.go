package solutions

func AppendCharacters(s string, t string) int {
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] == t[j] {
			j++
			if j == len(t) {
				return 0
			}
		}
	}
	return len(t) - j
}
