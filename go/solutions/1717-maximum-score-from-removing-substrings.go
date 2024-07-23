package solutions

func MaximumGain(s string, x int, y int) int {
	ans := 0

	str, val := s, max(x, y)
	first, last := byte('b'), byte('a')
	if x > y {
		first, last = last, first
	}

	for j := 0; j < 2; j++ {
		st := []byte{}
		for i := 0; i < len(str); i++ {
			if len(st) > 0 && str[i] == last && st[len(st)-1] == first {
				ans += val
				st = st[:len(st)-1]
			} else {
				st = append(st, str[i])
			}
		}
		str, val = string(st), min(x, y)
		first, last = last, first
	}

	return ans
}
