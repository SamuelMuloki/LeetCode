package solutions

func RemoveDuplicates4(s string, k int) string {
	st := []byte{}
	for i := range s {
		count := 1
		for ; len(st) >= count && s[i] == st[len(st)-count]; count++ {
		}

		if count == k {
			st = st[:len(st)-count+1]
		} else {
			st = append(st, s[i])
		}
	}

	return string(st)
}
