package solutions

import "container/heap"

type Boundary struct {
	Row, Col, Height int
}

func TrapRainWater(heightMap [][]int) int {
	dRow := []int{0, 0, -1, 1}
	dCol := []int{-1, 1, 0, 0}

	rows := len(heightMap)
	cols := len(heightMap[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var isValidCell = func(row, col int) bool {
		return row >= 0 && col >= 0 && row < rows && col < cols
	}

	boundaries := &BoundaryHeap{}
	heap.Init(boundaries)

	for i := 0; i < rows; i++ {
		heap.Push(boundaries, Boundary{i, 0, heightMap[i][0]})
		heap.Push(boundaries, Boundary{i, cols - 1, heightMap[i][cols-1]})
		visited[i][cols-1] = true
		visited[i][0] = visited[i][cols-1]
	}

	for i := 0; i < cols; i++ {
		heap.Push(boundaries, Boundary{0, i, heightMap[0][i]})
		heap.Push(boundaries, Boundary{rows - 1, i, heightMap[rows-1][i]})
		visited[rows-1][i] = true
		visited[0][i] = visited[rows-1][i]
	}

	ans := 0
	for boundaries.Len() > 0 {
		currentBoundary := heap.Pop(boundaries).(Boundary)
		currentRow, currentCol, minBoundaryHeight := currentBoundary.Row, currentBoundary.Col, currentBoundary.Height

		for direction := 0; direction < 4; direction++ {
			neighborRow := currentRow + dRow[direction]
			neighborCol := currentCol + dCol[direction]

			if isValidCell(neighborRow, neighborCol) && !visited[neighborRow][neighborCol] {
				neighborHeight := heightMap[neighborRow][neighborCol]

				if neighborHeight < minBoundaryHeight {
					ans += minBoundaryHeight - neighborHeight
				}

				heap.Push(boundaries, Boundary{neighborRow, neighborCol, max(neighborHeight, minBoundaryHeight)})
				visited[neighborRow][neighborCol] = true
			}
		}
	}

	return ans
}

type BoundaryHeap []Boundary

func (h BoundaryHeap) Len() int           { return len(h) }
func (h BoundaryHeap) Less(i, j int) bool { return h[i].Height < h[j].Height }
func (h BoundaryHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *BoundaryHeap) Push(x any)        { *h = append(*h, x.(Boundary)) }
func (h *BoundaryHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
