package solutions

func MaxWidthRamp(nums []int) int {
	n := len(nums)
	st := []int{}

	for i := 0; i < n; i++ {
		if len(st) == 0 || nums[st[len(st)-1]] > nums[i] {
			st = append(st, i)
		}
	}

	maxWidth := 0
	for j := n - 1; j >= 0; j-- {
		for len(st) > 0 && nums[st[len(st)-1]] <= nums[j] {
			i := st[len(st)-1]
			st = st[:len(st)-1]
			maxWidth = max(maxWidth, j-i)
		}
	}

	return maxWidth
}
