package solutions

import "strings"

func LargestGoodInteger(num string) string {
	ans := byte(0)
	for i := 2; i < len(num); i++ {
		if num[i] == num[i-1] && num[i] == num[i-2] {
			ans = max(ans, num[i])
		}
	}

	if ans == byte(0) {
		return ""
	}

	return strings.Repeat(string(ans), 3)
}
