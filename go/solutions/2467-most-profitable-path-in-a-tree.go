package solutions

import (
	"math"
)

func MostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(amount)
	maxIncome := math.MinInt32
	graph := make([][]int, n)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	visited := make([]bool, n)
	bobPath := make(map[int]int)
	var findBobPath func(sourceNode, time int) map[int]int
	findBobPath = func(sourceNode, time int) map[int]int {
		bobPath[sourceNode] = time
		visited[sourceNode] = true

		if sourceNode == 0 {
			return bobPath
		}

		for _, adjacentNode := range graph[sourceNode] {
			if !visited[adjacentNode] {
				if findBobPath(adjacentNode, time+1) != nil {
					return bobPath
				}
			}
		}

		delete(bobPath, sourceNode)
		return nil
	}

	bobPath = findBobPath(bob, 0)

	queue := [][]int{}
	queue = append(queue, []int{0, 0, 0})
	visited = make([]bool, n)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sourceNode, time, income := node[0], node[1], node[2]

		if _, found := bobPath[sourceNode]; !found || time < bobPath[sourceNode] {
			income += amount[sourceNode]
		} else if time == bobPath[sourceNode] {
			income += amount[sourceNode] / 2
		}

		if len(graph[sourceNode]) == 1 && sourceNode != 0 {
			maxIncome = max(maxIncome, income)
		}

		for _, adjacentNode := range graph[sourceNode] {
			if !visited[adjacentNode] {
				queue = append(queue, []int{adjacentNode, time + 1, income})
			}
		}

		visited[sourceNode] = true
	}

	return maxIncome
}
