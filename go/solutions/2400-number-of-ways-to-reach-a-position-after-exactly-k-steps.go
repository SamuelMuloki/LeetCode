package solutions

func NumberOfWays(startPos int, endPos int, k int) int {
	/*
		We'll consider the distance between start and end position since it's non negative
		hence easier to memoize, For k number of steps and distance diff,
		the number of ways is dfs(k - 1, diff + 1) + dfs(k - 1, abs(diff - 1))

		Note that for d == 0, the number of ways is dfs(k - 1, 1) + dfs(k - 1, 1), since abs(0 - 1) == 1.
		We can reach 0 from positions 1 and -1, and the number of ways for negative positions mirrors positive positions.
	*/
	dp := make([][]int, 1001)
	for i := range dp {
		dp[i] = make([]int, 1001)
	}

	for i := range dp {
		for k := range dp[i] {
			dp[i][k] = -1
		}
	}

	var abs = func(diff int) int {
		if diff < 0 {
			return -diff
		}

		return diff
	}

	var dfs func(diff, steps, ways int) int
	dfs = func(diff, steps, ways int) int {
		if diff == steps {
			return 1
		}

		if diff > steps {
			return 0
		}

		if dp[diff][steps] != -1 {
			return dp[diff][steps]
		}

		left := dfs(diff+1, steps-1, ways)
		right := dfs(abs(diff-1), steps-1, ways)
		dp[diff][steps] = (left + right) % (1e9 + 7)

		return dp[diff][steps]
	}

	return dfs(abs(startPos-endPos), k, 0)
}
