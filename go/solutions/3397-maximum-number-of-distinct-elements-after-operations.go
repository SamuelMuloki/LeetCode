package solutions

import (
	"sort"
)

func MaxDistinctElements(nums []int, k int) int {
	n := len(nums)
	intervals := make([][2]int64, 0, n)
	kk := int64(k)
	for _, x := range nums {
		xi := int64(x)
		intervals = append(intervals, [2]int64{
			xi - kk, xi + kk,
		})
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	curr := int64(-1 << 60)
	res := 0
	for _, interval := range intervals {
		l, r := interval[0], interval[1]
		var candidate int64 = max(curr+1, l)
		if candidate <= r {
			curr = candidate
			res++
		}
	}

	return res
}
