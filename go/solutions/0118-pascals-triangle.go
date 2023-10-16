// https://leetcode.com/problems/pascals-triangle/
package solutions

func Generate(numRows int) [][]int {
	output := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		output[i] = make([]int, i+1)
		output[i][0], output[i][i] = 1, 1
		for j := 1; j < i; j++ {
			output[i][j] = output[i-1][j-1] + output[i-1][j]
		}
	}

	return output
}
