package solutions

import "sort"

func DivideArray(nums []int, k int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return nil
		}

		ans = append(ans, []int{nums[i], nums[i+1], nums[i+2]})
	}

	return ans
}
