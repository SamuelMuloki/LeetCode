package solutions

import "sort"

func MaximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	ans := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] >= ans+1 {
			ans++
		}
	}

	return ans
}
