package solutions

import "sort"

func MaximumBeauty2(nums []int, k int) int {
	sort.Ints(nums)
	start := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]-nums[start] <= 2*k {
			ans++
			continue
		}
		start++
	}

	return ans
}
