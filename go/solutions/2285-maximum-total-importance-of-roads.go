package solutions

import "sort"

func MaximumImportance(n int, roads [][]int) int64 {
	degree := make([]int, n)
	for _, edge := range roads {
		degree[edge[0]]++
		degree[edge[1]]++
	}

	sort.Ints(degree)
	value, ans := int64(1), int64(0)
	for _, d := range degree {
		ans += (value * int64(d))
		value++
	}

	return ans
}
