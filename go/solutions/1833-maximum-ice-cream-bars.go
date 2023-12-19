package solutions

import "sort"

func MaxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	for i := range costs {
		if coins >= costs[i] {
			coins -= costs[i]
		} else {
			return i
		}
	}

	return len(costs)
}
