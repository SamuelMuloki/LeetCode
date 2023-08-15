package leetcode

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	subs := make([][]int, 0)
	curr := make([]int, 0)

	sort.Ints(nums)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		subs = append(subs, append([]int{}, curr...))
		if idx == len(nums) {
			return
		}

		for i := idx; i < len(nums); i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			curr = append(curr, nums[i])
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0)

	return subs
}
