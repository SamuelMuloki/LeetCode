package solutions

import (
	"sort"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func Insert(intervals [][]int, newInterval []int) [][]int {
	newIntervals := append([][]int{}, intervals...)
	j := sort.Search(len(newIntervals), func(k int) bool {
		return newIntervals[k][0] > newInterval[0]
	})

	if j > -1 {
		var fromJ [][]int
		if j < len(newIntervals) {
			fromJ = newIntervals[j:][:]
		}
		newIntervals = append([][]int{}, newIntervals[:j][:]...)
		newIntervals = append(newIntervals, newInterval)
		if len(fromJ) > 0 {
			newIntervals = append(newIntervals, fromJ...)
		}
	} else {
		newIntervals = append(newIntervals, newInterval)
	}

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
