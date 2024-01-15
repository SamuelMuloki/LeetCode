package solutions

import "sort"

func FindWinners(matches [][]int) [][]int {
	won, lost := make(map[int]int), make(map[int]int)
	for i := range matches {
		won[matches[i][0]]++
		lost[matches[i][1]]++
	}

	one, two := make([]int, 0), make([]int, 0)
	for k := range won {
		if _, ok := lost[k]; !ok {
			one = append(one, k)
		}
	}

	for k, v := range lost {
		if v == 1 {
			two = append(two, k)
		}
	}

	sort.Ints(one)
	sort.Ints(two)
	return [][]int{one, two}
}
