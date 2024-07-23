package solutions

import "math"

func JudgeSquareSum(c int) bool {
	for i := 0; i*i <= c; i++ {
		j := math.Sqrt(float64(c - i*i))
		if j == math.Trunc(j) {
			return true
		}
	}

	return false
}
