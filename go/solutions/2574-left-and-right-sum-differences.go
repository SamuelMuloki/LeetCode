package solutions

func LeftRightDifference(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	left[0] = 0
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i-1] + left[i-1]
	}

	right := make([]int, n)
	right[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		right[i] = nums[i+1] + right[i+1]
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = abs(left[i] - right[i])
	}

	return ans
}
