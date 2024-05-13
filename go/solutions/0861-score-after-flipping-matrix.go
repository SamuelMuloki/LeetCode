package solutions

func MatrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			for j := 0; j < n; j++ {
				grid[i][j] = 1 - grid[i][j]
			}
		}
	}

	for j := 1; j < n; j++ {
		countZero := 0
		for i := 0; i < m; i++ {
			if grid[i][j] == 0 {
				countZero++
			}
		}

		if countZero > m-countZero {
			for i := 0; i < m; i++ {
				grid[i][j] = 1 - grid[i][j]
			}
		}
	}

	score := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			columnScore := grid[i][j] << (n - j - 1)
			score += columnScore
		}
	}

	return score
}
