package solutions

import (
	"container/list"
	"math"
)

type MinimumCostEdge struct {
	node   int
	weight int
}

func MinimumCost2(n int, edges [][]int, query [][]int) []int {
	adjList := make([][]MinimumCostEdge, n)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], MinimumCostEdge{edge[1], edge[2]})
		adjList[edge[1]] = append(adjList[edge[1]], MinimumCostEdge{edge[0], edge[2]})
	}

	visited := make([]bool, n)
	components := make([]int, n)
	var componentCost []int

	var getComponentCost = func(source int, componentId int) int {
		nodesQueue := list.New()
		componentCost := math.MaxInt32

		nodesQueue.PushBack(source)
		visited[source] = true

		for nodesQueue.Len() > 0 {
			node := nodesQueue.Remove(nodesQueue.Front()).(int)
			components[node] = componentId

			for _, neighbor := range adjList[node] {
				componentCost &= neighbor.weight

				if visited[neighbor.node] {
					continue
				}
				visited[neighbor.node] = true
				nodesQueue.PushBack(neighbor.node)
			}
		}

		return componentCost
	}

	componentId := 0
	for node := 0; node < n; node++ {
		if !visited[node] {
			componentCost = append(componentCost, getComponentCost(node, componentId))
			componentId++
		}
	}

	var ans []int

	for _, q := range query {
		start := q[0]
		end := q[1]

		if components[start] == components[end] {
			ans = append(ans, componentCost[components[start]])
		} else {
			ans = append(ans, -1)
		}
	}

	return ans
}
