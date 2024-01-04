package solutions

import "sort"

func TargetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			break
		}

		if nums[i] != target {
			continue
		}

		ans = append(ans, i)
	}

	return ans
}
