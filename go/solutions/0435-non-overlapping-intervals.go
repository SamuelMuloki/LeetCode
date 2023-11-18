package solutions

import (
	"sort"
)

func EraseOverlapIntervals(intervals [][]int) int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res = append(res, intervals[0])
	count := 0
	for i := 1; i < len(intervals); i++ {
		n := len(res)
		if intervals[i][0] < res[n-1][1] {
			count++
			res[n-1][1] = min(res[n-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}

	return count
}
