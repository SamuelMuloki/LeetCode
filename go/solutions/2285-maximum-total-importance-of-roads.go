package solutions

import "sort"

func MaximumImportance(n int, roads [][]int) int64 {
	connections := make(map[int]int)
	for i := range roads {
		connections[roads[i][0]]++
		connections[roads[i][1]]++
	}

	keys := make([]int, 0)
	for k := range connections {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return connections[keys[i]] < connections[keys[j]]
	})

	assign, N := make(map[int]int), len(keys)
	start := n - N + 1
	for i := range keys {
		assign[keys[i]] = start + i
	}

	ans := 0
	for i := range roads {
		ans += assign[roads[i][0]] + assign[roads[i][1]]
	}

	return int64(ans)
}
