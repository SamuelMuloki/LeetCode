package solutions

func RemoveOccurrences(s string, part string) string {
	st, partLen := []byte{}, len(part)
	for i := range s {
		st = append(st, s[i])
		if len(st) >= partLen && string(st[len(st)-partLen:]) == part {
			st = st[:len(st)-partLen]
		}
	}

	return string(st)
}
