package solutions

func Rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}

		for k, n := 0, len(matrix[i])-1; k < n; k, n = k+1, n-1 {
			matrix[i][k], matrix[i][n] = matrix[i][n], matrix[i][k]
		}
	}
}
