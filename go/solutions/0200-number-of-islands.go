package solutions

func NumIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m {
			return
		}
		if j < 0 || j >= n {
			return
		}
		if grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i, row := range grid {
		for j, col := range row {
			if col == '1' {
				ans++
				dfs(i, j)
			}
		}
	}

	return ans
}
