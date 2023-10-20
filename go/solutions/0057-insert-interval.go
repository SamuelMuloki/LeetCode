package solutions

import (
	"sort"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func Insert(intervals [][]int, newInterval []int) [][]int {
	newIntervals := append([][]int{}, intervals...)
	newIntervals = append(newIntervals, newInterval)
	sort.Slice(newIntervals, func(i, j int) bool {
		return newIntervals[i][0] < newIntervals[j][0]
	})

	res := make([][]int, 0)
	res = append(res, newIntervals[0])
	for i := 1; i < len(newIntervals); i++ {
		n := len(res)
		if res[n-1][1] >= newIntervals[i][0] {
			res[n-1][1] = utils.Max(res[n-1][1], newIntervals[i][1])
		} else {
			res = append(res, newIntervals[i])
		}
	}

	return res
}
