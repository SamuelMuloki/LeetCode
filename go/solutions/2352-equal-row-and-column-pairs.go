package solutions

func EqualPairs(grid [][]int) int {
	var res int
	set := make(map[[200]int]int)

	arr := [200]int{}
	for i := 0; i < len(grid); i++ {
		copy(arr[:], grid[i])
		set[arr]++
	}

	for i := 0; i < len(grid[0]); i++ {
		arr := [200]int{}
		for j := 0; j < len(grid); j++ {
			arr[j] = grid[j][i]
		}

		res += set[arr]
	}

	return res
}
