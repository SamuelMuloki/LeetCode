package solutions

func EventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	inStack := make([]bool, n)
	visit := make([]bool, n)

	var dfs func(i int) bool
	dfs = func(node int) bool {
		if inStack[node] {
			return true
		}

		if visit[node] {
			return false
		}

		inStack[node] = true
		visit[node] = true
		for _, neighbor := range graph[node] {
			if dfs(neighbor) {
				return true
			}
		}

		inStack[node] = false
		return false
	}

	for i := 0; i < n; i++ {
		dfs(i)
	}

	ans := []int{}
	for i := 0; i < n; i++ {
		if !inStack[i] {
			ans = append(ans, i)
		}
	}

	return ans
}
