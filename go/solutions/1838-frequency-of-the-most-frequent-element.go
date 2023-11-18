package solutions

import "sort"

func MaxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	left, curr := 0, 0
	for right := 0; right < len(nums); right++ {
		target := nums[right]
		curr += target

		if (right-left+1)*target-curr > k {
			curr -= nums[left]
			left++
		}
	}

	return len(nums) - left
}
