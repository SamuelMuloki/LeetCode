package solutions

func LargestLocal(grid [][]int) [][]int {
	n := len(grid)
	var findMax = func(i, j int) int {
		maxElement := 0
		for x := i; x < i+3; x++ {
			for y := j; y < j+3; y++ {
				maxElement = max(maxElement, grid[x][y])
			}
		}

		return maxElement
	}

	maxLocal := make([][]int, n-2)
	for i := range maxLocal {
		maxLocal[i] = make([]int, n-2)
	}

	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			maxLocal[i][j] = findMax(i, j)
		}
	}

	return maxLocal
}
