package solutions

func FindDuplicates(nums []int) []int {
	ans := []int{}
	for i := range nums {
		idx := abs(nums[i]) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		} else {
			ans = append(ans, abs(nums[i]))
		}
	}

	return ans
}
