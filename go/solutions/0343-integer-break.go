// https://leetcode.com/problems/integer-break/
package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func IntegerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}

	dp := make([]int, n+1)
	var dfs func(num int) int
	dfs = func(num int) int {
		if num <= 3 {
			return num
		}

		if dp[num] != 0 {
			return dp[num]
		}

		res := num
		for i := 2; i < num; i++ {
			res = utils.Max(res, i*dfs(num-i))
		}

		dp[num] = res
		return dp[num]
	}

	return dfs(n)
}
