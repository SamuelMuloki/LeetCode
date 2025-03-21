package solutions

import (
	"container/heap"
	"math"
)

func MaximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	if grid[n-1][n-1] == 1 || grid[0][0] == 1 {
		return 0
	}

	dist := computeDistanceGrid(grid)
	pq := &DistanceGridMaxHeap{}
	heap.Init(pq)
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	heap.Push(pq, &DistanceGridNode{dist[0][0], 0, 0})
	vis[0][0] = true

	for pq.Len() > 0 {
		distanceGridNode := heap.Pop(pq).(*DistanceGridNode)
		ds, i, j := distanceGridNode.dist, distanceGridNode.i, distanceGridNode.j

		if i == n-1 && j == n-1 {
			return ds
		}

		for _, dir := range directions {
			newRow, newCol := i+dir[0], j+dir[1]
			if isValid(newRow, newCol, n) && grid[newRow][newCol] != 1 && !vis[newRow][newCol] {
				vis[newRow][newCol] = true
				ds1 := dist[newRow][newCol]
				heap.Push(pq, &DistanceGridNode{min(ds, ds1), newRow, newCol})
			}
		}
	}

	return 0
}

func isValid(x, y, n int) bool {
	return x >= 0 && x < n && y >= 0 && y < n
}

func computeDistanceGrid(grid [][]int) [][]int {
	n := len(grid)
	distGrid := make([][]int, n)
	for i := range distGrid {
		distGrid[i] = make([]int, n)
		for j := range distGrid[i] {
			distGrid[i][j] = math.MaxInt32
		}
	}

	q := [][2]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, [2]int{i, j})
				distGrid[i][j] = 0
			}
		}
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if isValid(newX, newY, n) && distGrid[newX][newY] == math.MaxInt32 {
				distGrid[newX][newY] = distGrid[x][y] + 1
				q = append(q, [2]int{newX, newY})
			}
		}
	}

	return distGrid
}

type DistanceGridNode struct {
	dist, i, j int
}

type DistanceGridMaxHeap []*DistanceGridNode

func (h DistanceGridMaxHeap) Len() int            { return len(h) }
func (h DistanceGridMaxHeap) Less(i, j int) bool  { return h[i].dist > h[j].dist }
func (h DistanceGridMaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *DistanceGridMaxHeap) Push(x interface{}) { *h = append(*h, x.(*DistanceGridNode)) }
func (h *DistanceGridMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
