package solutions

func MaxAdjacentDistance(nums []int) int {
	n := len(nums)
	ans := max(abs(nums[n-1]-nums[0]), abs(nums[0]-nums[n-1]))
	for i := 1; i < n; i++ {
		ans = max(ans, max(
			abs(nums[i-1]-nums[i]),
			abs(nums[i]-nums[i-1])),
		)
	}

	return ans
}
