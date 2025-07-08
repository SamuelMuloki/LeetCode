package solutions

func CountSubarrays3(nums []int) int {
	res := 0
	for i := 2; i < len(nums); i++ {
		sum := nums[i-2] + nums[i]
		if nums[i-1]/2 == sum && nums[i-1]%2 == 0 {
			res++
		}
	}

	return res
}
