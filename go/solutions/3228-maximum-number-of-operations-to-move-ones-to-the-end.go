package solutions

func MaxOperations2(s string) int {
	st := []byte{}
	found := false
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			found = true
			continue
		}

		if len(st) > 0 && found {
			res += len(st)
		}

		st = append(st, '1')
		found = false
	}

	if len(st) > 0 && found {
		res += len(st)
	}

	return res
}
