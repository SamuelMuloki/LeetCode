// https://leetcode.com/problems/pascals-triangle/
package solutions

func Generate(numRows int) [][]int {
	output := make([][]int, numRows)

	output[0] = append(output[0], 1)
	for i := 1; i < numRows; i++ {
		newArr := []int{1}
		for j := 1; j < len(output[i-1]); j++ {
			newArr = append(newArr, output[i-1][j]+output[i-1][j-1])
		}

		newArr = append(newArr, 1)
		output[i] = append(output[i], newArr...)
	}

	return output
}
