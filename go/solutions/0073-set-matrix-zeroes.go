package solutions

func SetZeroes(matrix [][]int) {
	rc, cr := make(map[int]int), make(map[int]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				cr[j] = i
				rc[i] = j
			}
		}
	}

	for r := range rc {
		for i := 0; i < len(matrix[r]); i++ {
			matrix[r][i] = 0
		}
	}

	for c := range cr {
		for i := 0; i < len(matrix); i++ {
			matrix[i][c] = 0
		}
	}
}
