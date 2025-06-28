package solutions

import "sort"

func MaxSubsequence(nums []int, k int) []int {
	n := len(nums)
	arr := make([]int, n)
	copy(arr, nums)
	sort.Ints(arr)
	set := make(map[int]int)
	for i := n - k; i < n; i++ {
		set[arr[i]]++
	}

	res := []int{}
	for i := 0; i < n; i++ {
		if _, ok := set[nums[i]]; ok {
			res = append(res, nums[i])
			set[nums[i]]--
			if set[nums[i]] == 0 {
				delete(set, nums[i])
			}
		}
	}

	return res
}
