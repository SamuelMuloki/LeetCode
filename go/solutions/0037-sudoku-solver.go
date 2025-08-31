package solutions

func SolveSudoku(board [][]byte) {
	var isValid = func(row, col int, num byte) bool {
		for i := 0; i < 9; i++ {
			if board[row][i] == num || board[i][col] == num {
				return false
			}
		}

		startRow := (row / 3) * 3
		startCol := (col / 3) * 3

		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] == num {
					return false
				}
			}
		}

		return true
	}

	var solve func() bool
	solve = func() bool {
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if board[row][col] == '.' {
					for num := byte('1'); num <= '9'; num++ {
						if isValid(row, col, num) {
							board[row][col] = num
							if solve() {
								return true
							}
							board[row][col] = '.'
						}
					}
					return false
				}
			}
		}

		return true
	}

	solve()
}
