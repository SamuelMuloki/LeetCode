package solutions

func MinimumDeletions(s string) int {
	st, deleteCount := []byte{}, 0
	for i := range s {
		if len(st) > 0 && st[len(st)-1] == 'b' && s[i] == 'a' {
			st = st[:len(st)-1]
			deleteCount++
		} else {
			st = append(st, s[i])
		}
	}

	return deleteCount
}
