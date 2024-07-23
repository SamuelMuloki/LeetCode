package solutions

func CountSubarrays(nums []int, k int) int64 {
	ans, maxElement := 0, 0
	start, maxElementsInWindow := 0, 0
	for i := range nums {
		maxElement = max(maxElement, nums[i])
	}

	for end := 0; end < len(nums); end++ {
		if nums[end] == maxElement {
			maxElementsInWindow++
		}
		for k == maxElementsInWindow {
			if nums[start] == maxElement {
				maxElementsInWindow--
			}
			start++
		}

		ans += start
	}

	return int64(ans)
}
