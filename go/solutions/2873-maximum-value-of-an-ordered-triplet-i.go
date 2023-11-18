package solutions

func MaximumTripletValue(nums []int) int64 {
	maxV, n := 0, len(nums)
	maxLeft := make([]int, n)
	maxRight := make([]int, n)

	maxRight[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], nums[i])
	}

	maxLeft[0] = nums[0]
	for i := 1; i < n; i++ {
		maxLeft[i] = max(maxLeft[i-1], nums[i])
	}

	for j := 1; j < n-1; j++ {
		maxV = max(maxV, (maxLeft[j-1]-nums[j])*maxRight[j+1])
	}

	return int64(maxV)
}
