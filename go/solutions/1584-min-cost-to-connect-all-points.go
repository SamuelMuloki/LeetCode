package solutions

import (
	"sort"
)

type Edge struct {
	Vertex1 int
	Vertex2 int
	Weight  int
}

func MinCostConnectPoints(points [][]int) int {
	parent := make([]int, len(points))
	rank := make([]int, len(points))

	for i := 0; i < len(points); i++ {
		parent[i] = -1
		rank[i] = 1
	}

	var find func(i int) int
	find = func(i int) int {
		if parent[i] == -1 {
			return i
		}

		parent[i] = find(parent[i])

		return parent[i]
	}

	union := func(x, y int) bool {
		s1, s2 := find(x), find(y)

		if s1 == s2 {
			return false
		}

		switch {
		case rank[s1] < rank[s2]:
			parent[s1] = s2
		case rank[s1] > rank[s2]:
			parent[s2] = s1
		default:
			parent[s2] = s1
			rank[s1]++

		}

		return true
	}

	abs := func(val int) int {
		if val < 0 {
			return -val
		}

		return val
	}

	getWeight := func(pointX, pointY []int) int {
		return abs(pointX[0]-pointY[0]) + abs(pointX[1]-pointY[1])
	}

	edges := make([]Edge, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{
				Vertex1: i,
				Vertex2: j,
				Weight:  getWeight(points[i], points[j]),
			})
		}
	}

	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	var output int
	connected := 0

	for i := 0; i < len(edges); i++ {
		if union(edges[i].Vertex1, edges[i].Vertex2) {
			connected++
			output += edges[i].Weight
		}

		if connected == len(edges)-1 {
			return output
		}
	}

	return output
}
