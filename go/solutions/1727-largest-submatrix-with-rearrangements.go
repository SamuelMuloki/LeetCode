package solutions

import "sort"

func LargestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	ans, prevRow := 0, make([]int, n)
	for row := 0; row < m; row++ {
		currRow := append([]int{}, matrix[row]...)
		for col := 0; col < n; col++ {
			if matrix[row][col] != 0 && row > 0 {
				currRow[col] += prevRow[col]
			}
		}

		sortedRow := append([]int{}, currRow...)
		sort.Slice(sortedRow, func(i, j int) bool { return sortedRow[i] > sortedRow[j] })
		for i := 0; i < n; i++ {
			ans = max(ans, sortedRow[i]*(i+1))
		}

		prevRow = currRow
	}

	return ans
}
