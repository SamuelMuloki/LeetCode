package solutions

func CountSquares(matrix [][]int) int {
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row+1)
	for i := range dp {
		dp[i] = make([]int, col+1)
	}

	ans := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 1 {
				dp[i+1][j+1] = min(min(dp[i][j+1], dp[i+1][j]), dp[i][j]) + 1
				ans += dp[i+1][j+1]
			}
		}
	}

	return ans
}
