package solutions

import "sort"

func MaxSum3(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	res := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i+1] == nums[i] {
			continue
		}

		res = max(res, res+nums[i])
	}

	return res
}
