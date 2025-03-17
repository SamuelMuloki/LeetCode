package solutions

import "container/list"

func MagnificentSets(n int, edges [][]int) int {
	adjList := make([][]int, n)
	for _, edge := range edges {
		adjList[edge[0]-1] = append(adjList[edge[0]-1], edge[1]-1)
		adjList[edge[1]-1] = append(adjList[edge[1]-1], edge[0]-1)
	}

	colors := make([]int, n)
	for i := range colors {
		colors[i] = -1
	}

	var isBipartite func(node int) bool
	isBipartite = func(node int) bool {
		for _, neighbor := range adjList[node] {
			if colors[neighbor] == colors[node] {
				return false
			}
			if colors[neighbor] != -1 {
				continue
			}
			colors[neighbor] = (colors[node] + 1) % 2
			if !isBipartite(neighbor) {
				return false
			}
		}
		return true
	}

	var getLongestShortestPath = func(srcNode, n int) int {
		nodesQueue := list.New()
		visited := make([]bool, n)

		nodesQueue.PushBack(srcNode)
		visited[srcNode] = true
		distance := 0

		for nodesQueue.Len() > 0 {
			numOfNodesInLayer := nodesQueue.Len()
			for i := 0; i < numOfNodesInLayer; i++ {
				currentNode := nodesQueue.Remove(nodesQueue.Front()).(int)
				for _, neighbor := range adjList[currentNode] {
					if visited[neighbor] {
						continue
					}
					visited[neighbor] = true
					nodesQueue.PushBack(neighbor)
				}
			}
			distance++
		}
		return distance
	}

	var getNumberOfGroupsForComponent func(node int, distances []int, visited []bool) int
	getNumberOfGroupsForComponent = func(node int, distances []int, visited []bool) int {
		maxNumberOfGroups := distances[node]
		visited[node] = true

		for _, neighbor := range adjList[node] {
			if visited[neighbor] {
				continue
			}
			maxNumberOfGroups = max(maxNumberOfGroups, getNumberOfGroupsForComponent(neighbor, distances, visited))
		}
		return maxNumberOfGroups
	}

	for node := 0; node < n; node++ {
		if colors[node] != -1 {
			continue
		}
		colors[node] = 0
		if !isBipartite(node) {
			return -1
		}
	}

	distances := make([]int, n)
	for node := 0; node < n; node++ {
		distances[node] = getLongestShortestPath(node, n)
	}

	maxNumberOfGroups := 0
	visited := make([]bool, n)
	for node := 0; node < n; node++ {
		if visited[node] {
			continue
		}
		maxNumberOfGroups += getNumberOfGroupsForComponent(node, distances, visited)
	}

	return maxNumberOfGroups
}
