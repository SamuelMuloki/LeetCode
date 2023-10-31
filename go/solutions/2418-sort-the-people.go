package solutions

import "sort"

type People struct {
	name   string
	height int
}

func SortPeople(names []string, heights []int) []string {
	arr := make([]People, len(names))
	for i := 0; i < len(names); i++ {
		arr[i] = People{
			name:   names[i],
			height: heights[i],
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].height > arr[j].height
	})

	ans := make([]string, len(names))
	for i := range arr {
		ans[i] = arr[i].name
	}

	return ans
}
