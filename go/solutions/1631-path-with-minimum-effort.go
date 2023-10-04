package solutions

import (
	"container/heap"
	"math"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func MinimumEffortPath(heights [][]int) int {
	rows, cols := len(heights), len(heights[0])
	adj := make([][]int, rows)
	for i := range adj {
		adj[i] = make([]int, cols)
		for j := range adj[i] {
			adj[i][j] = math.MaxInt32
		}
	}

	abs := func(val int) int {
		if val < 0 {
			return -val
		}

		return val
	}

	adj[0][0] = 0
	moves := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	edges := &EdgeMinHeap{Edge{0, 0, 0}}
	heap.Init(edges)

	for edges.Len() > 0 {
		edge := heap.Pop(edges).(Edge)
		effort, v1, v2 := edge.Weight, edge.Vertex1, edge.Vertex2
		if effort > adj[v1][v2] {
			continue
		}
		if v1 == rows-1 && v2 == cols-1 {
			return effort
		}
		for _, move := range moves {
			nV1, nV2 := v1+move[0], v2+move[1]
			if nV1 >= 0 && nV1 < rows && nV2 >= 0 && nV2 < cols {
				newEffort := utils.Max(effort, abs(heights[v1][v2]-heights[nV1][nV2]))
				if newEffort < adj[nV1][nV2] {
					adj[nV1][nV2] = newEffort
					heap.Push(edges, Edge{nV1, nV2, newEffort})
				}
			}
		}
	}
	return -1
}
