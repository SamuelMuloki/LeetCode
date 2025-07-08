package solutions

import (
	"container/heap"
	"math"
)

func MinTimeToReach2(moveTime [][]int) int {
	n, m := len(moveTime), len(moveTime[0])
	var isValid = func(x, y int) bool {
		return x >= 0 && y >= 0 && x < n && y < m
	}

	d := make([][]int, n)
	visited := make([][]bool, n)
	for i := range visited {
		d[i] = make([]int, m)
		visited[i] = make([]bool, m)
		for j := range d[i] {
			d[i][j] = math.MaxInt32
		}
	}
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	d[0][0] = 0
	queue := &MinHeap2{}
	heap.Push(queue, State2{0, 0, 0})
	for queue.Len() > 0 {
		curr := heap.Pop(queue).(State2)
		if visited[curr.x][curr.y] {
			continue
		}

		visited[curr.x][curr.y] = true
		for _, dir := range directions {
			nextX, nextY := curr.x+dir[0], curr.y+dir[1]
			if isValid(nextX, nextY) {
				dist := max(d[curr.x][curr.y], moveTime[nextX][nextY]) + (curr.x+curr.y)%2 + 1
				if d[nextX][nextY] > dist {
					d[nextX][nextY] = dist
					heap.Push(queue, State2{nextX, nextY, dist})
				}
			}
		}
	}

	return d[n-1][m-1]
}
