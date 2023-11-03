package solutions

func BuildArray2(nums []int) []int {
	n := len(nums)
	for j := range nums {
		nums[j] = nums[j] + n*(nums[nums[j]]%n)
	}

	for i := range nums {
		nums[i] = nums[i] / n
	}

	return nums
}
