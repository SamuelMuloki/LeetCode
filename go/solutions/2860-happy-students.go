package solutions

import "sort"

func CountWays(nums []int) int {
	sort.Ints(nums)
	count, selected, n := 0, 0, len(nums)

	if nums[0] != 0 { // Not Selecting AnyOne
		count = 1
	}

	for i := 0; i < n; i++ {
		selected++
		if selected > nums[i] { // No. of Selected Students is strictly greater than nums[i].
			if i+1 < n && selected < nums[i+1] { // Considering from (i+1) to n Students is not Selected.
				count++
			} else if i+1 == n { // Last Student
				count++
			}
		}
	}

	return count
}
