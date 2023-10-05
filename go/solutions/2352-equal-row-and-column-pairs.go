package solutions

import "fmt"

func EqualPairs(grid [][]int) int {
	var res int
	set := make(map[string]int)

	for i := 0; i < len(grid); i++ {
		set[fmt.Sprintf("%v", grid[i])]++
	}

	for i := 0; i < len(grid[0]); i++ {
		arr := []int{}
		for j := 0; j < len(grid); j++ {
			arr = append(arr, grid[j][i])
		}

		res += set[fmt.Sprintf("%v", arr)]
	}

	return res
}
