package solutions

func MinimumTime(n int, relations [][]int, time []int) int {
	graph := make(map[int][]int)
	for _, edge := range relations {
		x, y := edge[0]-1, edge[1]-1
		graph[x] = append(graph[x], y)
	}

	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(node int) int
	dfs = func(node int) int {
		if memo[node] != -1 {
			return memo[node]
		}

		if len(graph[node]) == 0 {
			return time[node]
		}

		ans := 0
		for _, neighbor := range graph[node] {
			ans = max(ans, dfs(neighbor))
		}

		memo[node] = time[node] + ans
		return memo[node]
	}

	ans := 0
	for node := 0; node < n; node++ {
		ans = max(ans, dfs(node))
	}

	return ans
}
