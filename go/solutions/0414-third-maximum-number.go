package solutions

import "sort"

func ThirdMax(nums []int) int {
	arr := append([]int{}, nums...)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	set := make(map[int]int)
	for i := range arr {
		set[arr[i]]++

		if len(set) == 3 {
			return arr[i]
		}
	}

	return arr[0]
}
