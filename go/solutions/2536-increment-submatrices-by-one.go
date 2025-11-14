package solutions

func RangeAddQueries(n int, queries [][]int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	for _, query := range queries {
		startRow, startCol := query[0], query[1]
		endRow, endCol := query[2], query[3]
		for i := startRow; i <= endRow; i++ {
			for j := startCol; j <= endCol; j++ {
				res[i][j] += 1
			}
		}
	}

	return res
}
