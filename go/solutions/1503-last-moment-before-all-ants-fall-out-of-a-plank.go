package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func GetLastMoment(n int, left []int, right []int) int {
	ans := 0
	for i := range left {
		ans = utils.Max(ans, left[i])
	}

	for i := range right {
		ans = utils.Max(ans, n-right[i])
	}

	return ans
}
