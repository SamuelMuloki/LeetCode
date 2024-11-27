package solutions

func ShortestDistanceAfterQueries(n int, queries [][]int) []int {
	var bfs = func(adjList [][]int) int {
		visited := make([]bool, n)
		nodeQueue := []int{}
		nodeQueue = append(nodeQueue, 0)
		visited[0] = true

		currentLayerNodeCount, nextLayerNodeCount := 1, 0
		layersExplored := 0

		for len(nodeQueue) > 0 {
			for i := 0; i < currentLayerNodeCount; i++ {
				currentNode := nodeQueue[0]
				nodeQueue = nodeQueue[1:]
				if currentNode == n-1 {
					return layersExplored
				}

				for _, neighbor := range adjList[currentNode] {
					if visited[neighbor] {
						continue
					}
					nodeQueue = append(nodeQueue, neighbor)
					nextLayerNodeCount++
					visited[neighbor] = true
				}
			}

			currentLayerNodeCount = nextLayerNodeCount
			nextLayerNodeCount = 0
			layersExplored++
		}

		return -1
	}

	ans := []int{}
	adjList := make([][]int, n)
	for i := 0; i < n-1; i++ {
		adjList[i] = append(adjList[i], i+1)
	}

	for _, road := range queries {
		u, v := road[0], road[1]
		adjList[u] = append(adjList[u], v)
		ans = append(ans, bfs(adjList))
	}

	return ans
}
