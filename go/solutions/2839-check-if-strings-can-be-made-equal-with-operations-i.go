package solutions

func CanBeEqual(s1 string, s2 string) bool {
	i, j := 0, 2
	for j < 4 {
		if s1[i] == s2[i] {
			i++
			j++
			continue
		}

		if s1[i] != s2[j] {
			return false
		}

		s2 = swap([]byte(s2), i, j)
		i++
		j++
	}

	return s1 == s2
}

func swap(s []byte, i, j int) string {
	s[i], s[j] = s[j], s[i]

	return string(s)
}
