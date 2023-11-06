package solutions

import (
	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func NumWays(steps int, arrLen int) int {
	arrLen = utils.Min(arrLen, steps)
	dp := make([]int, arrLen)
	prevDp := make([]int, arrLen)

	prevDp[0] = 1
	for remain := 1; remain <= steps; remain++ {
		dp = make([]int, arrLen)
		for curr := arrLen - 1; curr >= 0; curr-- {
			res := prevDp[curr]
			if curr > 0 {
				res = (res + prevDp[curr-1]) % (1e9 + 7)
			}

			if curr < arrLen-1 {
				res = (res + prevDp[curr+1]) % (1e9 + 7)
			}

			dp[curr] = res
		}
		prevDp = dp
	}

	return dp[0]
}
