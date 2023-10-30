package solutions

func DiagonalSum(mat [][]int) int {
	ans, n := 0, len(mat)
	for i := 0; i < n; i++ {
		// Add elements from primary diagonal.
		ans += mat[i][i]
		// Add elements from secondary diagonal.
		ans += mat[n-1-i][i]
	}

	// If n is odd, subtract the middle element as its added twice.
	if n&1 != 0 {
		ans -= mat[n/2][n/2]
	}

	return ans
}
