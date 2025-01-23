package solutions

func CountServers(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	ans := 0
	rowCount, colCount := make(map[int]int), make(map[int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowCount[i]++
				colCount[j]++
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (rowCount[i] > 1 || colCount[j] > 1) {
				ans++
			}
		}
	}

	return ans
}
