// https://leetcode.com/problems/integer-break/
package solutions

import (
	"math"
)

func IntegerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}

	if n%3 == 0 {
		return int(math.Pow(float64(3), float64(n/3)))
	}

	if n%3 == 1 {
		return int(math.Pow(float64(3), float64((n/3)-1))) * 4
	}

	return int(math.Pow(float64(3), float64(n/3))) * 2
}
