package solutions

import "sort"

func MinimizeMax(nums []int, p int) int {
	var countValidPairs = func(arr []int, threshold int) int {
		index, count := 0, 0
		for index < len(arr)-1 {
			if nums[index+1]-nums[index] <= threshold {
				count++
				index++
			}
			index++
		}

		return count
	}

	sort.Ints(nums)
	n := len(nums)
	left, right := 0, nums[n-1]-nums[0]
	for left < right {
		mid := left + (right-left)/2

		if countValidPairs(nums, mid) >= p {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
