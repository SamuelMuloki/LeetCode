package solutions

func NumMagicSquaresInside(grid [][]int) int {
	count := 0
	rows, cols := len(grid), len(grid[0])

	for i := 0; i <= rows-3; i++ {
		for j := 0; j <= cols-3; j++ {
			if isValidMagicSquare(grid, i, j) {
				count++
			}
		}
	}

	return count
}

func isValidMagicSquare(grid [][]int, startRow, startCol int) bool {
	numPresence := make([]bool, 10)

	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			num := grid[i][j]
			if num < 1 || num > 9 || numPresence[num] {
				return false
			}
			numPresence[num] = true
		}
	}

	targetSum := grid[startRow][startCol] + grid[startRow][startCol+1] + grid[startRow][startCol+2]
	for i := 0; i < 3; i++ {
		if getRowSum(grid, startRow+i, startCol) != targetSum ||
			getColSum(grid, startRow, startCol+i) != targetSum {
			return false
		}
	}

	diagonalSum1 := grid[startRow][startCol] + grid[startRow+1][startCol+1] + grid[startRow+2][startCol+2]
	diagonalSum2 := grid[startRow+2][startCol] + grid[startRow+1][startCol+1] + grid[startRow][startCol+2]
	return diagonalSum1 == targetSum && diagonalSum2 == targetSum
}

func getRowSum(grid [][]int, row, startCol int) int {
	return grid[row][startCol] + grid[row][startCol+1] + grid[row][startCol+2]
}

func getColSum(grid [][]int, startRow, col int) int {
	return grid[startRow][col] + grid[startRow+1][col] + grid[startRow+2][col]
}
