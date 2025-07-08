package solutions

import "sort"

func CountDays(days int, meetings [][]int) int {
	dayMap := make(map[int]int)
	prevDay := days
	for _, meeting := range meetings {
		prevDay = min(prevDay, meeting[0])
		dayMap[meeting[0]]++
		dayMap[meeting[1]+1]--
	}

	keys := []int{}
	for k := range dayMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	sum := 0
	freeDays := prevDay - 1
	for _, currDay := range keys {
		if sum == 0 {
			freeDays += currDay - prevDay
		}

		sum += dayMap[currDay]
		prevDay = currDay
	}

	freeDays += days - prevDay + 1

	return freeDays
}
