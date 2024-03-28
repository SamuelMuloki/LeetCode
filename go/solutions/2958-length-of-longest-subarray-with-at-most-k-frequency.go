package solutions

func MaxSubarrayLength(nums []int, k int) int {
	ans, end := 0, -1
	window := make(map[int]int)

	for start := 0; start < len(nums); start++ {
		window[nums[start]]++
		for window[nums[start]] > k {
			end++
			window[nums[end]]--
		}
		ans = max(ans, start-end)
	}

	return ans
}
