package solutions

func LargestIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(row, col, islandId int) int
	dfs = func(row, col, islandId int) int {
		if row < 0 || col < 0 || row >= m || col >= n || grid[row][col] != 1 {
			return 0
		}

		count := 1
		grid[row][col] = islandId
		for d := 0; d < 4; d++ {
			count += dfs(row+dirs[d][0], col+dirs[d][1], islandId)
		}

		return count
	}

	islandSizes := make(map[int]int)
	islandId := 2
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				islandSizes[islandId] = dfs(i, j, islandId)
				islandId++
			}
		}
	}

	if len(islandSizes) == 0 {
		return 1
	}

	if len(islandSizes) == 1 {
		islandId--
		if islandSizes[islandId] == m*n {
			return islandSizes[islandId]
		}

		return islandSizes[islandId] + 1
	}

	maxIslandSize := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				currIslandSize := 1
				neighboringIslands := make(map[int]struct{})

				for _, dir := range dirs {
					newRow, newCol := i+dir[0], j+dir[1]
					if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && grid[newRow][newCol] > 1 {
						neighboringIslands[grid[newRow][newCol]] = struct{}{}
					}
				}

				for id := range neighboringIslands {
					currIslandSize += islandSizes[id]
				}

				maxIslandSize = max(maxIslandSize, currIslandSize)
			}
		}
	}

	return maxIslandSize
}
