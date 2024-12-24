package solutions

func MaxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	graph := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	result := 0
	var dfs func(node, parent int) int
	dfs = func(node, parent int) int {
		sum := values[node]

		for _, neighbor := range graph[node] {
			if neighbor != parent {
				sum += dfs(neighbor, node)
			}
		}

		if sum%k == 0 {
			result++
			return 0
		}
		return sum
	}

	dfs(0, -1)

	return result
}
