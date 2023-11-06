package solutions

import (
	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func MinimumTotal(triangle [][]int) int {
	n := len(triangle)
	minPath := triangle[n-1]
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			minPath[j] = utils.Min(minPath[j], minPath[j+1]) + triangle[i][j]
		}
	}

	return minPath[0]
}
