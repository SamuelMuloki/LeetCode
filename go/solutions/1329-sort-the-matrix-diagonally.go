package solutions

import "sort"

func DiagonalSort(mat [][]int) [][]int {
	diagonal := make(map[int][]int)
	for i := range mat {
		for j := range mat[i] {
			diagonal[i-j] = append(diagonal[i-j], mat[i][j])
		}
	}

	for _, diag := range diagonal {
		sort.Ints(diag)
	}

	for i := range mat {
		for j := range mat[i] {
			mat[i][j] = diagonal[i-j][0]
			diagonal[i-j] = diagonal[i-j][1:]
		}
	}

	return mat
}
