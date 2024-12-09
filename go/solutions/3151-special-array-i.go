package solutions

func IsArraySpecial(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}

	for i := 1; i < n; i++ {
		if nums[i-1]%2 == nums[i]%2 {
			return false
		}
	}

	return true
}
