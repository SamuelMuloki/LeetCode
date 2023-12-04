package solutions

func MaxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rowMaxes, colMaxes := make([]int, n), make([]int, n)

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			rowMaxes[r] = max(rowMaxes[r], grid[r][c])
			colMaxes[c] = max(colMaxes[c], grid[r][c])
		}
	}

	ans := 0
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			ans += min(rowMaxes[r], colMaxes[c]) - grid[r][c]
		}
	}

	return ans
}
