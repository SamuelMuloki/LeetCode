package solutions

import "sort"

func LengthOfLIS(nums []int) int {
	sub := make([]int, 0)

	for i := range nums {
		if len(sub) == 0 || sub[len(sub)-1] < nums[i] {
			sub = append(sub, nums[i])
		} else {
			idx := sort.Search(len(sub), func(k int) bool { return sub[k] >= nums[i] })
			sub[idx] = nums[i]
		}
	}

	return len(sub)
}
