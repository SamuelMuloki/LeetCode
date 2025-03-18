package solutions

func LongestNiceSubarray(nums []int) int {
	maxLen := 1
	curr_bits := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		for curr_bits&nums[right] != 0 {
			curr_bits ^= nums[left]
			left++
		}

		curr_bits |= nums[right]
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
