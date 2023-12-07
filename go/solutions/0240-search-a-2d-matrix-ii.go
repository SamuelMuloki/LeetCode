package solutions

func SearchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	row, col := 0, n-1
	for col >= 0 && row <= m-1 {
		if target == matrix[row][col] {
			return true
		} else if target < matrix[row][col] {
			col--
		} else if target > matrix[row][col] {
			row++
		}
	}

	return false
}
