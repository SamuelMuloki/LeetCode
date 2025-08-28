package solutions

import "sort"

func SortMatrix(grid [][]int) [][]int {
	n := len(grid)
	for i := 0; i < n; i++ {
		curr := []int{}
		for j := 0; i+j < n; j++ {
			curr = append(curr, grid[i+j][j])
		}
		sort.Sort(sort.Reverse(sort.IntSlice(curr)))
		for j := 0; i+j < n; j++ {
			grid[i+j][j] = curr[j]
		}
	}

	for j := 1; j < n; j++ {
		curr := []int{}
		for i := 0; j+i < n; i++ {
			curr = append(curr, grid[i][j+i])
		}
		sort.Ints(curr)
		for i := 0; j+i < n; i++ {
			grid[i][j+i] = curr[i]
		}
	}

	return grid
}
