package solutions

func ClearDigits(s string) string {
	st := []byte{}
	for i := 0; i < len(s); i++ {
		if len(st) > 0 && s[i] >= '0' && s[i] <= '9' {
			st = st[:len(st)-1]
		} else {
			st = append(st, s[i])
		}
	}

	return string(st)
}
