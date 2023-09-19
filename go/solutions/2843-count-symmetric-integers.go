package solutions

import (
	"strconv"
)

func CountSymmetricIntegers(low int, high int) int {
	n := 0
	for i := low; i <= high; i++ {
		if !isSymmetric(i) {
			continue
		}

		n++
	}

	return n
}

func isSymmetric(n int) bool {
	str := strconv.Itoa(n)
	if len(str)%2 != 0 {
		return false
	}

	l, r := 0, 0
	for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
		l += int(str[i] - '0')
		r += int(str[j] - '0')
	}

	return l == r
}
