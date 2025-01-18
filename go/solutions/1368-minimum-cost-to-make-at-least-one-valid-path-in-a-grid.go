package solutions

import (
	"container/heap"
	"math"
)

func MinCost2(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	minCost := make([][]int, rows)
	for i := range minCost {
		minCost[i] = make([]int, cols)
		for j := range minCost[i] {
			minCost[i][j] = math.MaxInt
		}
	}

	minCost[0][0] = 0
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	edges := &EdgeMinHeap{Edge{0, 0, 0}}
	heap.Init(edges)
	for edges.Len() > 0 {
		edge := heap.Pop(edges).(Edge)
		cost, row, col := edge.Weight, edge.Vertex1, edge.Vertex2

		if minCost[row][col] != cost {
			continue
		}

		for dir := 0; dir < 4; dir++ {
			newRow, newCol := row+dirs[dir][0], col+dirs[dir][1]
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
				newCost := cost
				if dir != grid[row][col]-1 {
					newCost++
				}
				if minCost[newRow][newCol] > newCost {
					minCost[newRow][newCol] = newCost
					heap.Push(edges, Edge{newRow, newCol, newCost})
				}
			}
		}
	}

	return minCost[rows-1][cols-1]
}
