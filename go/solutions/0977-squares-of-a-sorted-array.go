package solutions

import "sort"

func SortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i] * nums[i]
	}

	sort.Ints(ans)
	return ans
}
