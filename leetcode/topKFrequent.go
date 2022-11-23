package leetcode

import (
	"sort"
)

/*
Given an integer array nums and an integer k, return the
k most frequent elements. You may return the answer in any order.

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Input: nums = [1], k = 1
Output: [1]
*/

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
