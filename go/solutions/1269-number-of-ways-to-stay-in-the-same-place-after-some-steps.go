package solutions

import (
	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func NumWays(steps int, arrLen int) int {
	arrLen = utils.Min(arrLen, steps)
	dp := make([][]int, arrLen)
	for i := range dp {
		dp[i] = make([]int, steps+1)
	}

	dp[0][0] = 1
	for remain := 1; remain <= steps; remain++ {
		for curr := arrLen - 1; curr >= 0; curr-- {
			res := dp[curr][remain-1]
			if curr > 0 {
				res = (res + dp[curr-1][remain-1]) % (1e9 + 7)
			}

			if curr < arrLen-1 {
				res = (res + dp[curr+1][remain-1]) % (1e9 + 7)
			}

			dp[curr][remain] = res
		}
	}

	return dp[0][steps]
}
