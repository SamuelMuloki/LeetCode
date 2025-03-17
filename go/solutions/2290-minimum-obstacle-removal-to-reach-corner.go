package solutions

import (
	"container/list"
	"math"
)

type PointState struct {
	x       int
	y       int
	removed int
}

func MinimumObstacles(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	q := list.New()
	q.PushBack(PointState{0, 0, 0})
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	directions := []PointState{
		{0, 1, 0},
		{1, 0, 0},
		{-1, 0, 0},
		{0, -1, 0},
	}
	minObstacles := math.MaxInt
	for q.Len() > 0 {
		curr := q.Front().Value.(PointState)
		q.Remove(q.Front())

		for _, dir := range directions {
			nextX := curr.x + dir.x
			nextY := curr.y + dir.y
			removed := curr.removed
			if nextX == rows-1 && nextY == cols-1 {
				minObstacles = min(minObstacles, curr.removed)
				continue
			}
			if nextX >= 0 && nextY >= 0 && nextX < rows && nextY < cols && !visited[nextX][nextY] {
				visited[nextX][nextY] = true

				if grid[nextX][nextY] == 1 {
					removed++
					nextState := PointState{nextX, nextY, removed}
					q.PushBack(nextState)
				} else {
					nextState := PointState{nextX, nextY, removed}
					q.PushFront(nextState)
				}

			}
		}
	}
	return minObstacles
}
