package solutions

import "sort"

func FindLucky(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	set := make(map[int]int)
	for _, num := range arr {
		set[num]++
	}

	for i := n - 1; i >= 0; i-- {
		if set[arr[i]] == arr[i] {
			return arr[i]
		}
	}

	return -1
}
