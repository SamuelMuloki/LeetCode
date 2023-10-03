package solutions

import "sort"

func NumIdenticalPairs(nums []int) int {
	sort.Ints(nums)

	count := 0
	for i := 0; i < len(nums); i++ {
		j := i + 1
		for j < len(nums) && nums[i] == nums[j] {
			count++
			j++
		}
	}

	return count
}
