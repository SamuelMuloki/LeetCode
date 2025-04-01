package solutions

func MostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)
	for i := range dp {
		dp[i] = -1
	}
	var dfs func(curr int) int64
	dfs = func(curr int) int64 {
		if curr >= n {
			return 0
		}

		if dp[curr] != -1 {
			return dp[curr]
		}

		take := int64(questions[curr][0]) + dfs(curr+questions[curr][1]+1)
		leave := dfs(curr + 1)

		dp[curr] = max(take, leave)

		return dp[curr]
	}

	return dfs(0)
}
