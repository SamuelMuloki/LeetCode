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

	var abs = func(diff int) int {
		if diff < 0 {
			return -diff
		}

		return diff
	}

	for k := 1; k <= 1000; k++ {
		dp[k][k] = 1
		for i := 0; i < k; i++ {
			dp[k][i] = (dp[k-1][i+1] + dp[k-1][abs(i-1)]) % (1e9 + 7)
		}
	}

	return dp[k][abs(startPos-endPos)]
}
