package solutions

func SwimInWater(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	low, high := grid[0][0], n*n-1
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	canReach := func(t int) bool {
		if grid[0][0] > t {
			return false
		}
		visited := make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, m)
		}
		visited[0][0] = true
		queue := [][]int{{0, 0}}
		for len(queue) > 0 {
			x, y := queue[0][0], queue[0][1]
			queue = queue[1:]
			if x == n-1 && y == m-1 {
				return true
			}
			for _, dir := range directions {
				nx, ny := x+dir[0], y+dir[1]
				if nx >= 0 && nx < n && ny >= 0 && ny < m && !visited[nx][ny] && grid[nx][ny] <= t {
					visited[nx][ny] = true
					queue = append(queue, []int{nx, ny})
				}
			}
		}
		return false
	}

	for low < high {
		mid := (low + high) / 2
		if canReach(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
