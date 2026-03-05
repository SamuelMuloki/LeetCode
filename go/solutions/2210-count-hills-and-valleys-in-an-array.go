package solutions

func CountHillValley(nums []int) int {
	j := 0
	res := 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != nums[i-1] {
			j = i - 1
		}
		if nums[i] > nums[j] && nums[i] > nums[i+1] ||
			nums[i] < nums[j] && nums[i] < nums[i+1] {
			res++
		}
	}

	return res
}
