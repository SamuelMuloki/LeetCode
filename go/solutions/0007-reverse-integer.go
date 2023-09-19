package solutions

import "math"

func Reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
	}

	reversed, temp := 0, x*sign
	for temp > 0 {
		reversed = (reversed * 10) + (temp % 10)

		if reversed >= math.MaxInt32 || reversed <= math.MinInt32 {
			return 0
		}

		temp /= 10
	}

	return reversed * sign
}
