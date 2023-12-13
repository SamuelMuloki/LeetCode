package solutions

func OnesMinusZeros(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	rowCount, colCount := make([]int, m), make([]int, n)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				rowCount[r]++
				colCount[c]++
			}
		}
	}

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			ans[r][c] = (rowCount[r] + colCount[c]) - (n - rowCount[r]) - (m - colCount[c])
		}
	}

	return ans
}
