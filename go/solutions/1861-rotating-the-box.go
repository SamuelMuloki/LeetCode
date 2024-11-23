package solutions

func RotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])
	result := make([][]byte, n)
	for i := 0; i < n; i++ {
		result[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			result[i][j] = '.'
		}
	}

	for i := 0; i < m; i++ {
		lowestRowWithEmptyCell := n - 1
		for j := n - 1; j >= 0; j-- {
			if box[i][j] == '#' {
				result[lowestRowWithEmptyCell][m-i-1] = '#'
				lowestRowWithEmptyCell--
			}

			if box[i][j] == '*' {
				result[j][m-i-1] = '*'
				lowestRowWithEmptyCell = j - 1
			}
		}
	}
	return result
}
