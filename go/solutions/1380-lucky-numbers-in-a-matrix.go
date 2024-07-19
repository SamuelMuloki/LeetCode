package solutions

import "math"

func LuckyNumbers(matrix [][]int) []int {
	ans := []int{}
	for i := 0; i < len(matrix); i++ {
		currMin, minIndex := math.MaxInt, 0
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < currMin {
				currMin = matrix[i][j]
				minIndex = j
			}
		}

		found := true
		for k := 0; k < len(matrix); k++ {
			if matrix[k][minIndex] > currMin {
				found = false
				break
			}
		}

		if found {
			ans = append(ans, currMin)
		}
	}

	return ans
}
