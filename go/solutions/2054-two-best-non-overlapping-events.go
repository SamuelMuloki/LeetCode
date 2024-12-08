package solutions

import "sort"

func MaxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	n := len(events)
	arr := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		arr[i] = max(arr[i+1], events[i][2])
	}

	ans := 0
	for _, event := range events {
		maxValue := event[2]
		left, right := 0, n
		for left < right {
			mid := (left + right) >> 1
			if events[mid][0] > event[1] {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if left < n {
			maxValue += arr[left]
		}
		ans = max(ans, maxValue)
	}

	return ans
}
