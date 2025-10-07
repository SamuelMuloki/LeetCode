package solutions

import "sort"

func AvoidFlood(rains []int) []int {
	n := len(rains)
	res := make([]int, n)
	lakeToDay := map[int]int{}
	dryDays := []int{}

	for i := 0; i < n; i++ {
		if rains[i] == 0 {
			dryDays = append(dryDays, i)
			res[i] = 1
		} else {
			lake := rains[i]
			if day, exists := lakeToDay[lake]; exists {
				idx := sort.Search(len(dryDays), func(i int) bool { return dryDays[i] > day })
				if idx == len(dryDays) {
					return []int{}
				}
				dryDay := dryDays[idx]
				res[dryDay] = lake
				dryDays = append(dryDays[:idx], dryDays[idx+1:]...)
			}
			lakeToDay[lake] = i
			res[i] = -1
		}
	}

	return res
}
