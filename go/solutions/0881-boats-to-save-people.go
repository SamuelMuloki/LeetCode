package solutions

import "sort"

func NumRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	min, max := 0, len(people)-1
	ans := 0

	for min <= max {
		if people[min]+people[max] <= limit {
			min += 1
		}
		max -= 1
		ans += 1
	}

	return ans
}
