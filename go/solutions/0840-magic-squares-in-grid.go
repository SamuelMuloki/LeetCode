package solutions

func NumMagicSquaresInside(grid [][]int) int {
	ans := 0
	rows, cols := len(grid), len(grid[0])

	for row := 0; row+2 < rows; row++ {
		for col := 0; col+2 < cols; col++ {
			if isValidMagicSquare(grid, row, col) {
				ans++
			}
		}
	}

	return ans
}

func isValidMagicSquare(grid [][]int, row, col int) bool {
	seen := make([]bool, 10)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := grid[row+i][col+j]
			if num < 1 || num > 9 || seen[num] {
				return false
			}
			seen[num] = true
		}
	}

	diagonal1 := grid[row][col] + grid[row+1][col+1] + grid[row+2][col+2]
	diagonal2 := grid[row+2][col] + grid[row+1][col+1] + grid[row][col+2]

	if diagonal1 != diagonal2 {
		return false
	}

	row1 := grid[row][col] + grid[row][col+1] + grid[row][col+2]
	row2 := grid[row+1][col] + grid[row+1][col+1] + grid[row+1][col+2]
	row3 := grid[row+2][col] + grid[row+2][col+1] + grid[row+2][col+2]

	if !(row1 == diagonal1 && row2 == diagonal1 && row3 == diagonal1) {
		return false
	}

	col1 := grid[row][col] + grid[row+1][col] + grid[row+2][col]
	col2 := grid[row][col+1] + grid[row+1][col+1] + grid[row+2][col+1]
	col3 := grid[row][col+2] + grid[row+1][col+2] + grid[row+2][col+2]

	return (col1 == diagonal1 && col2 == diagonal1 && col3 == diagonal1)
}
