package solutions

func IsValid2(s string) bool {
	st := []byte{}
	for i := range s {
		if s[i] == 'c' && len(st) > 1 && st[len(st)-1] == 'b' && st[len(st)-2] == 'a' {
			st = st[:len(st)-2]
		} else {
			st = append(st, s[i])
		}
	}

	return len(st) == 0
}
