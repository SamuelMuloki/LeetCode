package solutions

func CountSubIslands(grid1 [][]int, grid2 [][]int) int {
	numRows, numCols := len(grid2), len(grid2[0])
	totalCells := numRows * numCols
	islandRoot := make([]int, totalCells)
	islandStatus := make([]byte, totalCells)

	for i := 0; i < totalCells; i++ {
		islandRoot[i] = i
	}

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if grid2[row][col] == 1 {
				currentIndex := row*numCols + col
				if col+1 < numCols && grid2[row][col+1] == 1 {
					unionIslands(islandRoot, currentIndex, currentIndex+1)
				}
				if row+1 < numRows && grid2[row+1][col] == 1 {
					unionIslands(islandRoot, currentIndex, currentIndex+numCols)
				}
			}
		}
	}

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if grid2[row][col] == 1 && grid1[row][col] == 0 {
				rootIndex := findIslandRoot(islandRoot, row*numCols+col)
				islandStatus[rootIndex] = 2
			}
		}
	}

	subIslandCount := 0
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if grid2[row][col] == 1 {
				rootIndex := findIslandRoot(islandRoot, row*numCols+col)
				if islandStatus[rootIndex] == 0 {
					subIslandCount++
					islandStatus[rootIndex] = 1
				}
			}
		}
	}

	return subIslandCount
}

func findIslandRoot(islandRoot []int, x int) int {
	if islandRoot[x] != x {
		islandRoot[x] = findIslandRoot(islandRoot, islandRoot[x])
	}
	return islandRoot[x]
}

func unionIslands(islandRoot []int, x, y int) {
	rootX := findIslandRoot(islandRoot, x)
	rootY := findIslandRoot(islandRoot, y)
	if rootX != rootY {
		islandRoot[rootY] = rootX
	}
}
