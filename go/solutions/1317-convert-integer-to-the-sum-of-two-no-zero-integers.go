package solutions

import (
	"strconv"
	"strings"
)

func GetNoZeroIntegers(n int) []int {
	for i := 1; i <= n; i++ {
		str1 := strconv.Itoa(i)
		str2 := strconv.Itoa(n - i)
		if !strings.Contains(str1, "0") && !strings.Contains(str2, "0") {
			return []int{n - i, i}
		}
	}

	return []int{}
}
