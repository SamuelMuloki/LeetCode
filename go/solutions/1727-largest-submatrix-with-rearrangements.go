package solutions

import "sort"

func LargestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	ans := 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] != 0 && row > 0 {
				matrix[row][col] += matrix[row-1][col]
			}
		}

		currRow := append([]int{}, matrix[row]...)
		sort.Ints(currRow)
		for i := 0; i < n; i++ {
			ans = max(ans, currRow[i]*(n-i))
		}
	}

	return ans
}
