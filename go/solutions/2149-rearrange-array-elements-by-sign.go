package solutions

func RearrangeArray(nums []int) []int {
	ans := make([]int, len(nums))
	pos, neg := 0, 1
	for i := range nums {
		if nums[i] > 0 {
			ans[pos] = nums[i]
			pos += 2
		} else {
			ans[neg] = nums[i]
			neg += 2
		}
	}

	return ans
}
