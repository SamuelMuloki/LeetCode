package solutions

func Transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	res := make([][]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j] = append(res[j], matrix[i][j])
		}
	}
	return res
}
