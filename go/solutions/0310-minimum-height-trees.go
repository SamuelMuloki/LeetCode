package solutions

func FindMinHeightTrees(n int, edges [][]int) []int {
	if n < 2 {
		return []int{0}
	}

	var adj = make([][]int, n)
	for i := 0; i < len(edges); i++ {
		var src, des = edges[i][0], edges[i][1]

		adj[src] = append(adj[src], des)
		adj[des] = append(adj[des], src)
	}

	var degree = make([]int, n)
	for _, edge := range edges {
		degree[edge[0]]++
		degree[edge[1]]++
	}

	var queue []int
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}

	var remaining = n
	for remaining > 2 {
		var size = len(queue)
		remaining -= size
		for i := 0; i < size; i++ {
			var v = queue[0]
			queue = queue[1:]
			for _, c := range adj[v] {
				degree[c]--
				if degree[c] == 1 {
					queue = append(queue, c)
				}
			}
		}
	}

	return queue
}
