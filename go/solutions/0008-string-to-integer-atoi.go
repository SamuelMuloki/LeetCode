// https://leetcode.com/problems/string-to-integer-atoi/
package solutions

import (
	"math"
)

func MyAtoi(s string) int {
	i, sLen := 0, len(s)
	if sLen == 0 {
		return 0
	}

	for i < sLen && s[i] == ' ' {
		i++
	}

	pos, neg := 0, 0
	if i < sLen && s[i] == '+' {
		pos++
		i++
	}

	if i < sLen && s[i] == '-' {
		neg++
		i++
	}

	n := 0
	for i < sLen && s[i] >= '0' && s[i] <= '9' {
		n = n*10 + int(s[i]-'0')

		if neg > 0 && n >= math.MaxInt32 {
			n = -n
			neg--
		}

		if n >= math.MaxInt32 {
			return math.MaxInt32
		}

		if n <= math.MinInt32 {
			return math.MinInt32
		}

		i++
	}

	if neg > 0 {
		n = -n
	}

	if pos > 0 && neg > 0 {
		return 0
	}

	return n
}
