package solutions

import (
	"sort"
)

func MinPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums)/2; i++ {
		ans = max(ans, nums[i]+nums[len(nums)-1-i])
	}

	return ans
}
