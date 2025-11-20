package solutions

import "sort"

func IntersectionSizeTwo(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})

	todo := make([]int, n)
	for i := range todo {
		todo[i] = 2
	}

	ress := 0
	for t := n - 1; t >= 0; t-- {
		s := intervals[t][0]
		m := todo[t]
		for p := s; p < s+m; p++ {
			for i := 0; i <= t; i++ {
				if todo[i] > 0 && p <= intervals[i][1] {
					todo[i]--
				}
			}
			ress++
		}
	}

	return ress
}
