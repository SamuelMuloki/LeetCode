package solutions

func BackspaceCompare(s string, t string) bool {
	st1, st2 := make([]byte, 0), make([]byte, 0)
	for i := range s {
		if s[i] == '#' {
			if len(st1) > 0 {
				st1 = st1[:len(st1)-1]
			}
		} else {
			st1 = append(st1, s[i])
		}
	}

	for i := range t {
		if t[i] == '#' {
			if len(st2) > 0 {
				st2 = st2[:len(st2)-1]
			}
		} else {
			st2 = append(st2, t[i])
		}
	}

	return string(st1) == string(st2)
}
