package solutions

func FindDisappearedNumbers(nums []int) []int {
	ans := make([]int, 0)
	for i := range nums {
		idx := abs(nums[i]) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	for i := 1; i <= len(nums); i++ {
		if nums[i-1] > 0 {
			ans = append(ans, i)
		}
	}

	return ans
}
