package solutions

import "strconv"

func IsStrictlyPalindromic(n int) bool {
	for base := 2; base <= n-2; base++ {
		str := strconv.FormatInt(int64(n), base)
		for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
			if str[i] != str[j] {
				return false
			}
		}
	}

	return true
}
