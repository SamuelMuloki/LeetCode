package solutions

func RemoveDuplicateLetters(s string) string {
	st := make([]rune, 0)
	seen := make(map[rune]bool)
	lastOcc := make(map[rune]int)

	for i, c := range s {
		lastOcc[c] = i
	}

	for i, c := range s {
		if !seen[c] {
			for len(st) > 0 && c < st[len(st)-1] && i < lastOcc[st[len(st)-1]] {
				delete(seen, st[len(st)-1])
				st = st[:len(st)-1]
			}
			seen[c] = true
			st = append(st, c)
		}
	}

	return string(st)
}
