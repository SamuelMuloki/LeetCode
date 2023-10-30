package solutions

func DiagonalSum(mat [][]int) int {
	sum := 0
	for i := 0; i < len(mat); i++ {
		sum += mat[i][i]
		mat[i][i] = 0
	}

	for i, j := 0, len(mat)-1; i < len(mat); i, j = i+1, j-1 {
		sum += mat[i][j]
	}

	return sum
}
