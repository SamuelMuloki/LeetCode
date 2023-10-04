package solutions

func TotalNQueens(n int) int {
	count := 0
	// posDiag r + c, negDiag r - c
	col, posDiag, negDiag := make(map[int]int), make(map[int]int), make(map[int]int)

	var backtrack func(c int)
	backtrack = func(c int) {
		if c == n {
			count++
			return
		}

		for r := 0; r < n; r++ {
			if col[r] > 0 || posDiag[r+c] > 0 || negDiag[r-c] > 0 {
				continue
			}

			col[r], posDiag[r+c], negDiag[r-c] = 1, 1, 1
			backtrack(c + 1)
			col[r], posDiag[r+c], negDiag[r-c] = 0, 0, 0
		}
	}

	backtrack(0)

	return count
}
