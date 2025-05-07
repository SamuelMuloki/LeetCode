package solutions

import "container/list"

type Point3 struct {
	x, y, removed, cnt int
}

func ShortestPath(grid [][]int, k int) int {
	n, m := len(grid), len(grid[0])
	queue := list.New()
	queue.PushBack(Point3{0, 0, 0, 0})
	visited := make([][][]bool, n)
	for i := range visited {
		visited[i] = make([][]bool, m)
		for j := range visited[i] {
			visited[i][j] = make([]bool, k+1)
		}
	}
	visited[0][0][0] = true
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	var isValid = func(x, y, removed int) bool {
		return x >= 0 && y >= 0 && x < n && y < m &&
			removed <= k && !visited[x][y][removed]
	}

	for queue.Len() > 0 {
		curr := queue.Front().Value.(Point3)
		queue.Remove(queue.Front())
		if curr.x == n-1 && curr.y == m-1 {
			return curr.cnt
		}

		for _, dir := range directions {
			nextX, nextY := curr.x+dir[0], curr.y+dir[1]
			removed := curr.removed

			if isValid(nextX, nextY, removed) {
				if grid[nextX][nextY] == 1 {
					removed++
				}

				if removed <= k && !visited[nextX][nextY][removed] {
					visited[nextX][nextY][removed] = true
					queue.PushBack(Point3{nextX, nextY, removed, curr.cnt + 1})
				}
			}
		}
	}

	return -1
}
