package solutions

func IslandPerimeter(grid [][]int) int {
	islands, neighbors := 0, 0
	row, col := len(grid), len(grid[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				islands++
				if i < row-1 && grid[i+1][j] == 1 {
					neighbors++
				}

				if j < col-1 && grid[i][j+1] == 1 {
					neighbors++
				}
			}
		}
	}

	return islands*4 - neighbors*2
}
