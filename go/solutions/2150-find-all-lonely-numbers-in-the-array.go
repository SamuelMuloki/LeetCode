package solutions

import "sort"

func FindLonely(nums []int) []int {
	sort.Ints(nums)
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] || i < len(nums)-1 && nums[i+1] == nums[i] {
			continue
		}

		if i > 0 && nums[i-1]+1 == nums[i] || i < len(nums)-1 && nums[i+1]-1 == nums[i] {
			continue
		}

		ans = append(ans, nums[i])
	}

	return ans
}
