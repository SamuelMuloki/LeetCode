package solutions

import "math"

func GetLengthOfOptimalCompression(s string, k int) int {
	memo := make([][]int, len(s)+1)
	for i := range memo {
		memo[i] = make([]int, k+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(s string, k int) int
	dp = func(s string, k int) int {
		if memo[len(s)][k] != -1 {
			return memo[len(s)][k]
		}

		if len(s) <= k {
			return 0
		}

		ans := math.MaxInt32
		if k > 0 {
			ans = dp(s[1:], k-1)
		}

		count, currentK := 0, k
		for i := range s {
			if s[i] == s[0] {
				count++
			} else if currentK > 0 {
				currentK--
			} else {
				break
			}

			ans = min(ans, getLength(count)+dp(s[i+1:], currentK))
		}

		memo[len(s)][k] = ans
		return ans
	}

	return dp(s, k)
}

func getLength(x int) int {
	if x == 1 {
		return 1
	} else if x <= 9 {
		return 2
	} else if x <= 99 {
		return 3
	}

	return 4
}
