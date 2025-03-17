package solutions

import "math"

func MaxMatrixSum(matrix [][]int) int64 {
	totalSum := 0
	minAbsVal, negativeCount := math.MaxInt, 0
	for _, row := range matrix {
		for _, col := range row {
			totalSum += abs(col)
			if col < 0 {
				negativeCount++
			}
			minAbsVal = min(minAbsVal, abs(col))
		}
	}

	if negativeCount%2 != 0 {
		totalSum -= 2 * minAbsVal
	}

	return int64(totalSum)
}
