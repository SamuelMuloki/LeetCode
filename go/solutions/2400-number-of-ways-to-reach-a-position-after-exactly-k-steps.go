package solutions

func NumberOfWays(startPos int, endPos int, k int) int {
	/*
		1 <= startPos, endPos, k <= 1000
		We ONLY see three ranges of 1000 for start ranging from [ -1000 to 2000 ]
		So only 3000 max states can be there for start
	*/
	dp := make([][]int, 3020)
	for i := range dp {
		dp[i] = make([]int, 1005)
	}

	for i := range dp {
		for k := range dp[i] {
			dp[i][k] = -1
		}
	}

	var dfs func(start, steps, ways int) int
	dfs = func(start, steps, ways int) int {
		if start == endPos && steps == 0 {
			return 1
		}

		if steps == 0 {
			return 0
		}

		if dp[start+1000][steps] != -1 {
			return dp[start+1000][steps]
		}

		left := dfs(start+1, steps-1, ways)
		right := dfs(start-1, steps-1, ways)
		dp[start+1000][steps] = (left + right) % (1e9 + 7)

		return dp[start+1000][steps]
	}

	return dfs(startPos, k, 0)
}
