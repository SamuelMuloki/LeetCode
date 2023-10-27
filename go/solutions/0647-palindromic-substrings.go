package solutions

func CountSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += expand2(i, i, s)
		count += expand2(i, i+1, s)
	}

	return count
}

func expand2(i, j int, s string) int {
	l, r, count := i, j, 0

	for l >= 0 && r < len(s) && s[l] == s[r] {
		count++
		l--
		r++
	}

	return count
}
