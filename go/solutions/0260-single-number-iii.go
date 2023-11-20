package solutions

func SingleNumber3(nums []int) []int {
	diff := 0
	for i := range nums {
		diff ^= nums[i]
	}

	diff &= -diff
	ans := []int{0, 0}
	for i := range nums {
		if nums[i]&diff == 0 {
			ans[0] ^= nums[i]
		} else {
			ans[1] ^= nums[i]
		}
	}

	return ans
}
