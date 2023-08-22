package solutions

import (
	"sort"
)

func TopKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)

	for i := range nums {
		frequency[nums[i]] += 1
	}

	nonRepeat := []int{}
	for num := range frequency {
		nonRepeat = append(nonRepeat, num)
	}

	sort.SliceStable(nonRepeat, func(a, b int) bool {
		return frequency[nonRepeat[a]] > frequency[nonRepeat[b]]
	})

	return nonRepeat[:k]
}
