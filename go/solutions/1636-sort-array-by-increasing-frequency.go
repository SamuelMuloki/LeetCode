package solutions

import "sort"

func FrequencySort2(nums []int) []int {
	count := make(map[int]int)
	for i := range nums {
		count[nums[i]]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if count[nums[i]] == count[nums[j]] {
			return nums[i] > nums[j]
		}

		return count[nums[i]] < count[nums[j]]
	})

	return nums
}
