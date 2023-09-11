package solutions

func GroupThePeople(groupSizes []int) [][]int {
	output := make([][]int, 0)
	sizes := make(map[int][]int)

	for i := 0; i < len(groupSizes); i++ {
		sizes[groupSizes[i]] = append(sizes[groupSizes[i]], i)
	}

	for n, arr := range sizes {
		k := n
		for i := 0; i < len(arr); i += n {
			output = append(output, arr[i:k])
			k += n
		}
	}

	return output
}
