package solutions

func SmallestNumber3(pattern string) string {
	n := len(pattern)
	st := []int{}
	ans := []byte{}
	for i := 0; i <= n; i++ {
		st = append(st, i+1)
		if i == n || pattern[i] == 'I' {
			for len(st) > 0 {
				ans = append(ans, byte(st[len(st)-1]+'0'))
				st = st[:len(st)-1]
			}
		}
	}

	return string(ans)
}
