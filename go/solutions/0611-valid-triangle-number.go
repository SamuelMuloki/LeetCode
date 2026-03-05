package solutions

import "sort"

func TriangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	res := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			left, right := j+1, n
			target := nums[i] + nums[j]
			for left < right {
				mid := left + (right-left)/2
				if nums[mid] < target {
					left = mid + 1
				} else {
					right = mid
				}
			}
			res += left - (j + 1)
		}
	}

	return res
}
