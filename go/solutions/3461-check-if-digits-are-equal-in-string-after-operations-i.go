package solutions

func HasSameDigits(s string) bool {
	st := []rune(s)
	for len(st) > 2 {
		curr := make([]rune, 0)
		for j := 1; j < len(st); j++ {
			num := ((st[j-1] + st[j]) - '0') % 10
			curr = append(curr, rune(num+'0'))
		}
		st = curr
	}

	return st[0] == st[1]
}
