package solutions

import "sort"

func SmallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	arr := make([]int, n)
	copy(arr, nums)
	sort.Ints(arr)
	set := make(map[int]int)
	for i := range arr {
		if _, ok := set[arr[i]]; !ok {
			set[arr[i]] = i
		}
	}

	ans := make([]int, n)
	for i := range nums {
		ans[i] = set[nums[i]]
	}

	return ans
}
