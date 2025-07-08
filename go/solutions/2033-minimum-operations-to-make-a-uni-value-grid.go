package solutions

import "sort"

func MinOperations8(grid [][]int, x int) int {
	merged := make([]int, 0)
	for _, row := range grid {
		merged = append(merged, row...)
	}

	sort.Ints(merged)
	n := len(merged)
	ans := 0
	for i := 0; i < len(merged); i++ {
		diff := abs(merged[n/2] - merged[i])
		if diff%x != 0 {
			return -1
		}
		ans += diff / x
	}

	return ans
}
