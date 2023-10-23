package solutions

func RearrangeArray(nums []int) []int {
	ans := make([]int, len(nums))
	j := 0
	for i := range nums {
		if nums[i] > 0 {
			ans[j] = nums[i]
			j += 2
		}
	}

	j = 1
	for i := range nums {
		if nums[i] < 0 {
			ans[j] = nums[i]
			j += 2
		}
	}

	return ans
}
