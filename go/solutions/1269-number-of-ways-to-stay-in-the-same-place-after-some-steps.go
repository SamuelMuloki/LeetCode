package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func NumWays(steps int, arrLen int) int {
	arrLen = utils.Min(arrLen, steps)
	dp := make([][]int, arrLen)
	for i := range dp {
		dp[i] = make([]int, steps+1)
	}

	for i := range dp {
		for k := range dp[i] {
			dp[i][k] = -1
		}
	}

	var dfs func(curr, remain int) int
	dfs = func(curr, remain int) int {
		if remain == 0 {
			if curr == 0 {
				return 1
			}
			return 0
		}

		if dp[curr][remain] != -1 {
			return dp[curr][remain]
		}

		res := dfs(curr, remain-1)
		if curr > 0 {
			res = (res + dfs(curr-1, remain-1)) % (1e9 + 7)
		}

		if curr < arrLen-1 {
			res = (res + dfs(curr+1, remain-1)) % (1e9 + 7)
		}

		dp[curr][remain] = res
		return res
	}

	return dfs(0, steps)
}
