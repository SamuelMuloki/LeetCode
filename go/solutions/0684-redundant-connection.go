package solutions

func FindRedundantConnection(edges [][]int) []int {
	n := len(edges)
	adj := make([][]int, n+1)

	var dfs func(src, target int, visited []bool) bool
	dfs = func(src, target int, visited []bool) bool {
		visited[src] = true
		if src == target {
			return true
		}

		isFound := false
		for _, nei := range adj[src] {
			if !visited[nei] {
				isFound = isFound || dfs(nei, target, visited)
			}
		}

		return isFound
	}

	for _, edge := range edges {
		visited := make([]bool, n+1)
		if dfs(edge[0], edge[1], visited) {
			return edge
		}
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	return []int{}
}
