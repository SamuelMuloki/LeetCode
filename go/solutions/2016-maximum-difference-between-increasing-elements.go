package solutions

func MaximumDifference(nums []int) int {
	res := -1
	n := len(nums)
	maxRight := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], nums[i+1])
	}

	for i := 0; i < n; i++ {
		if nums[i] < maxRight[i] {
			res = max(res, maxRight[i]-nums[i])
		}
	}

	return res
}
