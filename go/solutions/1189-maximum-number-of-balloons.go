package solutions

func MaxNumberOfBalloons(text string) int {
	b, a, l, o, n := 0, 0, 0, 0, 0
	for _, r := range text {
		switch r {
		case 'b':
			b++
		case 'a':
			a++
		case 'l':
			l++
		case 'o':
			o++
		case 'n':
			n++
		}
	}

	ans := min(b, min(a, n))
	ans = min(ans, min(l/2, o/2))
	return ans
}
