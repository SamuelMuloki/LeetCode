package solutions

import "sort"

func FindPairs(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	ans := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			diff := nums[i] - nums[j]
			if diff == k {
				ans++
				break
			}

			if diff > k {
				break
			}
		}
	}

	return ans
}
