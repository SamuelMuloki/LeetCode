package solutions

import "sort"

func CanMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != arr[1]-arr[0] {
			return false
		}
	}

	return true
}
