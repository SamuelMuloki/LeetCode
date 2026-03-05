package solutions

func IsValidSudoku(board [][]byte) bool {
	dp := [9][9]bool{}
	dpRow := [9][9]bool{}
	dpCol := [9][9]bool{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				subBoxIndex := (i/3)*3 + (j / 3)
				if dp[subBoxIndex][num] || dpRow[i][num] || dpCol[j][num] {
					return false
				}
				dpRow[i][num] = true
				dpCol[j][num] = true
				dp[subBoxIndex][num] = true
			}
		}
	}

	return true
}
