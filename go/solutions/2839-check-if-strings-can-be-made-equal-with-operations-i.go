package solutions

func CanBeEqual(s1 string, s2 string) bool {
	t1, t2, t3 := []byte(s1), []byte(s1), []byte(s1)
	swap(t1, 0, 2)
	swap(t2, 1, 3)
	swap(t3, 0, 2)
	swap(t3, 1, 3)

	return s1 == s2 || string(t1[:]) == s2 || string(t2[:]) == s2 || string(t3[:]) == s2
}

func swap(s []byte, i, j int) string {
	s[i], s[j] = s[j], s[i]

	return string(s)
}
