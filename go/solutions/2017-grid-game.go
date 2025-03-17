package solutions

import "math"

func GridGame(grid [][]int) int64 {
	firstRowSum := int64(0)
	for _, val := range grid[0] {
		firstRowSum += int64(val)
	}
	secondRowSum := int64(0)
	minimumSum := int64(math.MaxInt64)
	for turnIndex := 0; turnIndex < len(grid[0]); turnIndex++ {
		firstRowSum -= int64(grid[0][turnIndex])
		minimumSum = min(minimumSum, max(firstRowSum, secondRowSum))
		secondRowSum += int64(grid[1][turnIndex])
	}
	return minimumSum
}
