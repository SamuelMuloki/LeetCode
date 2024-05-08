package solutions

import (
	"sort"
	"strconv"
)

func FindRelativeRanks(score []int) []string {
	sorted := make([]int, len(score))
	copy(sorted, score)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	set := make(map[int]string)
	for i := range sorted {
		if i == 0 {
			set[sorted[i]] = "Gold Medal"
		} else if i == 1 {
			set[sorted[i]] = "Silver Medal"
		} else if i == 2 {
			set[sorted[i]] = "Bronze Medal"
		} else {
			set[sorted[i]] = strconv.Itoa(i + 1)
		}
	}

	ans := make([]string, len(score))
	for i := range score {
		ans[i] = set[score[i]]
	}

	return ans
}
