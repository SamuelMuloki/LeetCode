package solutions

func NumSpecial(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	rowCount, colCount := make([]int, m), make([]int, n)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if mat[r][c] == 1 {
				rowCount[r]++
				colCount[c]++
			}
		}
	}

	ans := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if mat[r][c] == 1 {
				if rowCount[r] == 1 && colCount[c] == 1 {
					ans++
				}
			}
		}
	}
	return ans
}
