package solutions

import (
	"container/list"
	"math"
)

func MostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(amount)
	maxIncome := math.MinInt32
	tree := make([][]int, n)
	nodeQueue := list.New()
	nodeQueue.PushBack([]int{0, 0, 0})

	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}

	var findBobPath func(tree [][]int, sourceNode, time int, bobPath map[int]int, visited []bool) map[int]int
	findBobPath = func(tree [][]int, sourceNode, time int, bobPath map[int]int, visited []bool) map[int]int {
		bobPath[sourceNode] = time
		visited[sourceNode] = true

		if sourceNode == 0 {
			return bobPath
		}

		for _, adjacentNode := range tree[sourceNode] {
			if !visited[adjacentNode] {
				if findBobPath(tree, adjacentNode, time+1, bobPath, visited) != nil {
					return bobPath
				}
			}
		}

		delete(bobPath, sourceNode)
		return nil
	}

	bobPath := findBobPath(tree, bob, 0, make(map[int]int), make([]bool, n))

	visited := make([]bool, n)
	for nodeQueue.Len() > 0 {
		front := nodeQueue.Front()
		nodeQueue.Remove(front)
		node := front.Value.([]int)
		sourceNode, time, income := node[0], node[1], node[2]

		if _, found := bobPath[sourceNode]; !found || time < bobPath[sourceNode] {
			income += amount[sourceNode]
		} else if time == bobPath[sourceNode] {
			income += amount[sourceNode] / 2
		}

		if len(tree[sourceNode]) == 1 && sourceNode != 0 {
			if income > maxIncome {
				maxIncome = income
			}
		}

		for _, adjacentNode := range tree[sourceNode] {
			if !visited[adjacentNode] {
				nodeQueue.PushBack([]int{adjacentNode, time + 1, income})
			}
		}

		visited[sourceNode] = true
	}
	return maxIncome
}
