package solutions

func RemoveKdigits(num string, k int) string {
	st := []rune{}
	for _, ch := range num {
		for len(st) > 0 && st[len(st)-1] > ch && k > 0 {
			st = st[:len(st)-1]
			k--
		}
		st = append(st, ch)
	}

	for k > 0 {
		st = st[:len(st)-1]
		k--
	}

	for i := 0; i < len(st); i++ {
		if st[i] != '0' {
			return string(st[i:])
		}
	}

	return "0"
}
