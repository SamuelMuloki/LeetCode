package solutions

func HighestPeak(isWater [][]int) [][]int {
	rows, cols := len(isWater), len(isWater[0])
	dir := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
		for j := range grid[i] {
			grid[i][j] = -1
		}
	}

	queue := []Cell{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if isWater[i][j] == 1 {
				queue = append(queue, Cell{i, j})
				grid[i][j] = 0
			}
		}
	}

	heightOfNextLayer := 1
	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			for d := 0; d < 4; d++ {
				nRow, nCol := queue[i].row+dir[d][0], queue[i].col+dir[d][1]

				if nRow < rows && nRow >= 0 && nCol < cols && nCol >= 0 &&
					grid[nRow][nCol] == -1 {
					grid[nRow][nCol] = heightOfNextLayer

					queue = append(queue, Cell{nRow, nCol})
				}
			}
		}

		heightOfNextLayer++
		queue = queue[qLen:]
	}

	return grid
}
