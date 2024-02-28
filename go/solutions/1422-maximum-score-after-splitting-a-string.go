package solutions

import "math"

func MaxScore(s string) int {
	n := len(s)
	zeroes, ones, best := 0, 0, math.MinInt
	for i := 0; i < n-1; i++ {
		zeroes += int(s[i]-'0') ^ 1
		ones += int(s[i] - '0')

		best = max(best, zeroes-ones)
	}

	ones += int(s[n-1] - '0')
	return best + ones
}
