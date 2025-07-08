package solutions

func LargestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	adj := make([][]int, n)
	inDegree := make([]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		inDegree[v]++
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 26)
	}

	queue := []int{}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
		dp[i][colors[i]-'a'] = 1
	}

	visited := 0
	res := 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		visited++
		for _, v := range adj[u] {
			for c := 0; c < 26; c++ {
				if dp[v][c] < dp[u][c]+btoi(int(colors[v]-'a') == c) {
					dp[v][c] = dp[u][c] + btoi(int(colors[v]-'a') == c)
				}
			}
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
		for c := 0; c < 26; c++ {
			if dp[u][c] > res {
				res = dp[u][c]
			}
		}
	}

	if visited != n {
		return -1
	}
	return res
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
