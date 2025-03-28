package solutions

import (
	"container/heap"
	"sort"
)

type Point2 struct {
	x, y, num int
}

func MaxPoints2(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	sortedQueries := make([][2]int, len(queries))
	for i, query := range queries {
		sortedQueries[i] = [2]int{query, i}
	}

	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i][0] < sortedQueries[j][0]
	})

	h := &MaxPoints2MinHeap{}
	heap.Init(h)
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	count := 0
	heap.Push(h, Point2{0, 0, grid[0][0]})
	visited[0][0] = true

	ans := make([]int, len(queries))
	for _, query := range sortedQueries {
		queryValue, queryIndex := query[0], query[1]
		for h.Len() > 0 && (*h)[0].num < queryValue {
			pt := heap.Pop(h).(Point2)
			count++
			for _, dir := range dirs {
				newX, newY := pt.x+dir[0], pt.y+dir[1]
				if newX >= 0 && newX < m && newY >= 0 && newY < n && !visited[newX][newY] {
					visited[newX][newY] = true
					heap.Push(h, Point2{newX, newY, grid[newX][newY]})
				}
			}
		}
		ans[queryIndex] = count
	}

	return ans
}

type MaxPoints2MinHeap []Point2

func (h MaxPoints2MinHeap) Len() int           { return len(h) }
func (h MaxPoints2MinHeap) Less(i, j int) bool { return h[i].num < h[j].num }
func (h MaxPoints2MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxPoints2MinHeap) Push(x any)        { *h = append(*h, x.(Point2)) }
func (h *MaxPoints2MinHeap) Pop() any {
	old := *h
	n := len(old)
	tmp := old[n-1]
	*h = old[:n-1]
	return tmp
}
