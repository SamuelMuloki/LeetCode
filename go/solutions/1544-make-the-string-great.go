package solutions

func MakeGood(s string) string {
	if len(s) <= 1 {
		return s
	}

	st := []byte{}
	for i := 0; i < len(s); i++ {
		if len(st) > 0 {
			if s[i] >= 'a' && s[i] <= 'z' && s[i]-32 == st[len(st)-1] ||
				s[i] >= 'A' && s[i] <= 'Z' && s[i]+32 == st[len(st)-1] {
				st = st[:len(st)-1]
				continue
			}
		}

		st = append(st, s[i])
	}

	return string(st)
}
