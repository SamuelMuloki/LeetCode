package solutions

import "container/list"

func FindMaxFish(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	countFish := func(row, col int) int {
		count := 0
		cells := list.New()
		cells.PushBack(Cell{row, col})
		vis[row][col] = true
		for cells.Len() > 0 {
			cell := cells.Remove(cells.Front()).(Cell)
			row, col := cell.row, cell.col
			count += grid[row][col]

			for dir := 0; dir < 4; dir++ {
				nRow, nCol := row+dirs[dir][0], col+dirs[dir][1]
				if nRow < m && nRow >= 0 && nCol < n && nCol >= 0 &&
					grid[nRow][nCol] != 0 && !vis[nRow][nCol] {
					cells.PushBack(Cell{nRow, nCol})
					vis[nRow][nCol] = true
				}
			}
		}

		return count
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 && !vis[i][j] {
				ans = max(ans, countFish(i, j))
			}
		}
	}

	return ans
}
