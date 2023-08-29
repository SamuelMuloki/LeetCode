package solutions

func SolveNQueens(n int) [][]string {
	output, curr := make([][]string, 0), make([]string, 0)
	// posDiag r + c, negDiag r - c
	col, posDiag, negDiag := make(map[int]int), make(map[int]int), make(map[int]int)

	var backtrack func(c int)
	backtrack = func(c int) {
		if c == n {
			output = append(output, append([]string{}, curr...))
		}

		for r := 0; r < n; r++ {
			if col[r] > 0 || posDiag[r+c] > 0 || negDiag[r-c] > 0 {
				continue
			}

			col[r], posDiag[r+c], negDiag[r-c] = 1, 1, 1
			s := ""
			for i := 0; i < n; i++ {
				if i == r {
					s += "Q"
				} else {
					s += "."
				}
			}

			curr = append(curr, s)
			backtrack(c + 1)
			col[r], posDiag[r+c], negDiag[r-c] = 0, 0, 0
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0)
	return output
}
