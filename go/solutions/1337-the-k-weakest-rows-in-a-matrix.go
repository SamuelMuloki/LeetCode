package solutions

import (
	"sort"
)

func KWeakestRows(mat [][]int, k int) []int {
	soldiers := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		count := 0
		for j := 0; j < len(mat[i]); j++ {
			count += mat[i][j]
		}
		soldiers = append(soldiers, []int{i, count})
	}

	sort.SliceStable(soldiers[:], func(i, j int) bool {
		return soldiers[i][1] < soldiers[j][1]
	})

	count := make([]int, k)
	for i := 0; i < k; i++ {
		count[i] = soldiers[i][0]
	}

	return count
}
