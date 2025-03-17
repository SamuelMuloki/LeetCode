package solutions

func CountUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	const (
		UNGUARDED = 0
		GUARDED   = 1
		GUARD     = 2
		WALL      = 3
	)

	for _, guard := range guards {
		grid[guard[0]][guard[1]] = GUARD
	}

	for _, wall := range walls {
		grid[wall[0]][wall[1]] = WALL
	}

	var markGuarded = func(row, col int) {
		for r := row - 1; r >= 0; r-- {
			if grid[r][col] == GUARD || grid[r][col] == WALL {
				break
			}
			grid[r][col] = GUARDED
		}

		for r := row + 1; r < len(grid); r++ {
			if grid[r][col] == GUARD || grid[r][col] == WALL {
				break
			}
			grid[r][col] = GUARDED
		}

		for c := col - 1; c >= 0; c-- {
			if grid[row][c] == GUARD || grid[row][c] == WALL {
				break
			}
			grid[row][c] = GUARDED
		}

		for c := col + 1; c < len(grid[0]); c++ {
			if grid[row][c] == GUARD || grid[row][c] == WALL {
				break
			}
			grid[row][c] = GUARDED
		}
	}

	for _, guard := range guards {
		markGuarded(guard[0], guard[1])
	}

	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == UNGUARDED {
				count++
			}
		}
	}

	return count
}
