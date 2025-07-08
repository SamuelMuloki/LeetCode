package solutions

import "sort"

func PartitionArray(nums []int, k int) int {
	sort.Ints(nums)
	res := 1
	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-curr > k {
			res++
			curr = nums[i]
		}
	}

	return res
}
