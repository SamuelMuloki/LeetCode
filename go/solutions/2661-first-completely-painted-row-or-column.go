package solutions

type Cell struct {
	row, col int
}

func FirstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	count := make(map[int]Cell)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count[mat[i][j]] = Cell{i, j}
		}
	}

	rows, cols := make([]int, m), make([]int, n)
	for i := 0; i < len(arr); i++ {
		if cell, ok := count[arr[i]]; ok {
			rows[cell.row]++
			cols[cell.col]++

			if rows[cell.row] == n || cols[cell.col] == m {
				return i
			}
		}
	}

	return -1
}
