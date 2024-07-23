package solutions

func GetMaximumGold(grid [][]int) int {
	directions := []int{0, 1, 0, -1, 0}
	rows, cols := len(grid), len(grid[0])

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row < 0 || col < 0 || row == rows || col == cols || grid[row][col] == 0 {
			return 0
		}

		maxGold := 0
		originalVal := grid[row][col]
		grid[row][col] = 0
		for direction := 0; direction < 4; direction++ {
			maxGold = max(maxGold, dfs(directions[direction]+row, directions[direction+1]+col))
		}

		grid[row][col] = originalVal
		return maxGold + originalVal
	}

	ans := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			ans = max(ans, dfs(row, col))
		}
	}

	return ans
}
