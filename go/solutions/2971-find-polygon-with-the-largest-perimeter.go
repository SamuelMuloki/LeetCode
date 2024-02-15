package solutions

import "sort"

func LargestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	ans, prevSum := -1, 0
	for i := 0; i < len(nums); i++ {
		if prevSum > nums[i] {
			ans = prevSum + nums[i]
		}

		prevSum += nums[i]
	}

	return int64(ans)
}
