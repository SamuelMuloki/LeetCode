package solutions

func FindLexSmallestString(s string, a int, b int) string {
	n := len(s)
	res := s
	s = s + s
	g := gcd(b, n)
	for i := 0; i < n; i += g {
		for j := 0; j < 10; j++ {
			kLimit := 0
			if b%2 != 0 {
				kLimit = 9
			}
			for k := 0; k <= kLimit; k++ {
				t := []byte(s[i : i+n])
				for p := 1; p < n; p += 2 {
					t[p] = byte('0' + (int(t[p]-'0')+j*a)%10)
				}
				for p := 0; p < n; p += 2 {
					t[p] = byte('0' + (int(t[p]-'0')+k*a)%10)
				}
				tStr := string(t)
				if tStr < res {
					res = tStr
				}
			}
		}
	}
	return res
}
