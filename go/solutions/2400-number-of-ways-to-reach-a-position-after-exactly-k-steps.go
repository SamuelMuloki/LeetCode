package solutions

func NumberOfWays(startPos int, endPos int, k int) int {
	/*
		We'll consider the distance between start and end position since it's non negative
		hence easier to memoize, For k number of steps and distance diff,
		the number of ways is dfs(k - 1, diff + 1) + dfs(k - 1, abs(diff - 1))

		Note that for d == 0, the number of ways is dfs(k - 1, 1) + dfs(k - 1, 1), since abs(0 - 1) == 1.
		We can reach 0 from positions 1 and -1, and the number of ways for negative positions mirrors positive positions.
	*/
	dp := make([]int, 1001)
	prevDp := make([]int, 1001)
	var abs = func(diff int) int {
		if diff < 0 {
			return -diff
		}

		return diff
	}

	for j := 1; j <= k; j++ {
		dp = make([]int, 1001)
		dp[j] = 1
		for i := 0; i < j; i++ {
			dp[i] = (prevDp[i+1] + prevDp[abs(i-1)]) % (1e9 + 7)
		}
		prevDp = dp
	}

	return dp[abs(startPos-endPos)]
}
