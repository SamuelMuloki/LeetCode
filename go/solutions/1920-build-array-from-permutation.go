package solutions

func BuildArray2(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}

	return ans
}
