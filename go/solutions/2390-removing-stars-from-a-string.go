package solutions

func RemoveStars(s string) string {
	st := []rune{}
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if len(st) > 0 && runes[i] == '*' {
			st = st[:len(st)-1]
		} else {
			st = append(st, runes[i])
		}
	}

	return string(st)
}
