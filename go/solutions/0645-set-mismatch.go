package solutions

func FindErrorNums(nums []int) []int {
	var rep int
	for i := range nums {
		idx := abs(nums[i]) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		} else {
			rep = abs(nums[i])
		}
	}

	for i := range nums {
		if nums[i] > 0 {
			return []int{rep, i + 1}
		}
	}

	return nil
}
