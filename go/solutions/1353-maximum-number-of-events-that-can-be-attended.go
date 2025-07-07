package solutions

import "sort"

func MaxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	maxDay := 0
	for _, e := range events {
		if e[1] > maxDay {
			maxDay = e[1]
		}
	}

	parent := make([]int, maxDay+2)
	for i := 0; i <= maxDay+1; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	res := 0
	for _, e := range events {
		day := find(e[0])
		if day <= e[1] {
			res++
			parent[day] = find(day + 1)
		}
	}
	return res
}
