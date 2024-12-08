package solutions

import "sort"

func MaxTwoEvents(events [][]int) int {
	times := make([][3]int, 0)
	for _, e := range events {
		times = append(times, [3]int{e[0], 1, e[2]})
		times = append(times, [3]int{e[1] + 1, 0, e[2]})
	}

	ans, maxValue := 0, 0
	sort.Slice(times, func(i, j int) bool {
		if times[i][0] == times[j][0] {
			return times[i][1] < times[j][1]
		}
		return times[i][0] < times[j][0]
	})

	for _, time := range times {
		if time[1] == 1 {
			ans = max(ans, time[2]+maxValue)
		} else {
			maxValue = max(maxValue, time[2])
		}
	}

	return ans
}
