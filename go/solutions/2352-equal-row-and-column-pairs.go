package solutions

import (
	"strconv"
)

func EqualPairs(grid [][]int) int {
	var res int
	set := make(map[string]int)

	for i := 0; i < len(grid); i++ {
		str := ""
		for j := 0; j < len(grid[i]); j++ {
			str += strconv.Itoa(grid[i][j])
			str += "*"
		}
		set[str]++
	}

	for i := 0; i < len(grid[0]); i++ {
		str := ""
		for j := 0; j < len(grid); j++ {
			str += strconv.Itoa(grid[j][i])
			str += "*"
		}

		res += set[str]
	}

	return res
}
