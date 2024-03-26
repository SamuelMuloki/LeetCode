package solutions

func FirstMissingPositive(nums []int) int {
	n := len(nums)
	contains1 := false

	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			contains1 = true
		}

		if nums[i] <= 0 || nums[i] > n {
			nums[i] = 1
		}
	}

	if !contains1 {
		return 1
	}

	for i := 0; i < n; i++ {
		val := abs(nums[i])
		if val == n {
			nums[0] = -abs(nums[0])
		} else {
			nums[val] = -abs(nums[val])
		}
	}

	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return n
	}

	return n + 1
}
