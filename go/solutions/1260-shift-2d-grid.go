package solutions

func ShiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range grid {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			nI := (i + (j+k)/n) % m // (i + numbers of rows added)%m
			nJ := (j + k) % n       // (j + numbers of columns added)%n
			ans[nI][nJ] = grid[i][j]
		}
	}

	return ans
}
