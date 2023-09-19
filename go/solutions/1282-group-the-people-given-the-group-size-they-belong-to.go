package solutions

func GroupThePeople(groupSizes []int) [][]int {
	output := make([][]int, 0)
	sizes := make(map[int][]int)

	for i := 0; i < len(groupSizes); i++ {
		sizes[groupSizes[i]] = append(sizes[groupSizes[i]], i)

		if len(sizes[groupSizes[i]]) == groupSizes[i] {
			output = append(output, sizes[groupSizes[i]])
			sizes[groupSizes[i]] = []int{}
		}
	}

	return output
}
