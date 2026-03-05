package solutions

import "sort"

func MinimumBoxes(apple []int, capacity []int) int {
	sort.Ints(capacity)
	sum := 0
	for i := 0; i < len(apple); i++ {
		sum += apple[i]
	}

	res := 0
	for i := len(capacity) - 1; i >= 0; i-- {
		if sum <= 0 {
			break
		}

		sum -= capacity[i]
		res++
	}

	return res
}
