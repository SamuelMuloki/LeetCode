package solutions

func MaxAscendingSum(nums []int) int {
	ans := nums[0]
	curr := ans
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			curr = 0
		}

		curr += nums[i]
		ans = max(ans, curr)
	}

	return ans
}
