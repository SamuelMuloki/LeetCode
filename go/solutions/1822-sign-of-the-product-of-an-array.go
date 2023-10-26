package solutions

func ArraySign(nums []int) int {
	sign := 1
	for i := range nums {
		if nums[i] == 0 {
			return 0
		}

		if nums[i] < 0 {
			sign = -sign
		}
	}

	return sign
}
