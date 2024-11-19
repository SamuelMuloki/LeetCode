package solutions

func MinSteps3(n int) int {
	ans, d := 0, 2
	for n > 1 {
		for n%d == 0 {
			ans += d
			n /= d
		}
		d++
	}
	return ans
}
