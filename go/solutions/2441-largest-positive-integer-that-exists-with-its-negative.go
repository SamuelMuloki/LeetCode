package solutions

func FindMaxK(nums []int) int {
	set := make(map[int]bool)
	for i := range nums {
		set[nums[i]] = true
	}

	ans := -1
	for i := range nums {
		if nums[i] > 0 && set[-nums[i]] {
			ans = max(ans, nums[i])
		}
	}

	return ans
}
