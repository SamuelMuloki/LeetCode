package solutions

import "sort"

func FindLHS(nums []int) int {
	sort.Ints(nums)
	res := 0
	i := 0
	for j := 1; j < len(nums); j++ {
		for i < j && nums[j]-nums[i] > 1 {
			i++
		}
		if nums[i] == nums[j] {
			continue
		}
		res = max(res, j-i+1)
	}

	return res
}
