package solutions

import "sort"

func MaximumBeauty(items [][]int, queries []int) []int {
	result := make([]int, len(queries))
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	maxBeauty := make([]int, len(items))
	maxBeauty[0] = items[0][1]
	for i := 1; i < len(items); i++ {
		maxBeauty[i] = max(maxBeauty[i-1], items[i][1])
	}

	for i, query := range queries {
		index := sort.Search(len(items), func(j int) bool {
			return items[j][0] > query
		})

		if index > 0 {
			result[i] = maxBeauty[index-1]
		}
	}

	return result
}
