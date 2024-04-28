package solutions

func SumOfDistancesInTree(n int, edges [][]int) []int {
	ans := make([]int, n)
	if len(edges) >= n {
		return ans
	}

	count := make([]int, n)
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	var postOrder func(int, int)
	postOrder = func(src, parent int) {
		for _, adj := range graph[src] {
			if adj == parent {
				continue
			}
			postOrder(adj, src)
			count[src] += count[adj]
			ans[src] += ans[adj] + count[adj]
		}
		count[src] += 1
	}

	var preOrder func(int, int)
	preOrder = func(src, parent int) {
		for _, adj := range graph[src] {
			if parent == adj {
				continue
			}
			ans[adj] = ans[src] - count[adj] + n - count[adj]
			preOrder(adj, src)
		}
	}

	postOrder(0, -1)
	preOrder(0, -1)

	return ans
}
