package solutions

func MinimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	var bfs = func(graph map[int][]int, start int) (int, int) {
		dist := make(map[int]int)
		queue := []int{start}
		dist[start] = 0
		var farthestNode, maxDist int

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			for _, neighbor := range graph[current] {
				if _, visited := dist[neighbor]; !visited {
					dist[neighbor] = dist[current] + 1
					queue = append(queue, neighbor)
					if dist[neighbor] > maxDist {
						maxDist = dist[neighbor]
						farthestNode = neighbor
					}
				}
			}
		}

		return farthestNode, maxDist
	}

	var treeDiameter = func(edges [][]int) int {
		graph := make(map[int][]int)
		for _, edge := range edges {
			graph[edge[0]] = append(graph[edge[0]], edge[1])
			graph[edge[1]] = append(graph[edge[1]], edge[0])
		}

		farthestNode, _ := bfs(graph, 0)
		_, diameter := bfs(graph, farthestNode)

		return diameter
	}

	diameter1 := treeDiameter(edges1)
	diameter2 := treeDiameter(edges2)

	ans := max(diameter1, diameter2)
	ans = max(ans, (diameter1+1)/2+(diameter2+1)/2+1)

	return ans
}
