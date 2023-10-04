package solutions

func ChampagneTower(poured int, query_row int, query_glass int) float64 {
	mat := [101][101]float64{}

	mat[0][0] = float64(poured)
	for i := 0; i <= query_row; i++ {
		for j := 0; j <= i; j++ {
			if mat[i][j] > 1 {
				mat[i+1][j] += (mat[i][j] - 1) / 2
				mat[i+1][j+1] += (mat[i][j] - 1) / 2
				mat[i][j] = 1
			}
		}
	}

	return mat[query_row][query_glass]
}
