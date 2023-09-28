package solutions

import (
	"math"
	"strings"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := math.MaxInt
	for i := range strs {
		minLen = utils.Min(minLen, len(strs[i]))
	}

	l, r := 1, minLen
	for l <= r {
		mid := (l + r) / 2
		if isCommonPrefix(strs, mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return strs[0][:(l+r)/2]
}

func isCommonPrefix(strs []string, mid int) bool {
	str1 := strs[0][:mid]
	for i := 1; i < len(strs); i++ {
		if !strings.HasPrefix(strs[i], str1) {
			return false
		}
	}

	return true
}
