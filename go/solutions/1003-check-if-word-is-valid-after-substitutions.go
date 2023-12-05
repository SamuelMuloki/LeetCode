package solutions

func IsValid2(s string) bool {
	st := []byte{}
	for i := range s {
		st = append(st, s[i])
		if len(st) > 2 && string(st[len(st)-3:]) == "abc" {
			st = st[:len(st)-3]
		}
	}

	return len(st) == 0
}
