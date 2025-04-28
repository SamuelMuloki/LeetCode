package solutions

func CountSubarrays4(nums []int, k int64) int64 {
	n := len(nums)
	var res int64
	var currSum int64
	start := 0
	for end := 0; end < n; end++ {
		currSum += int64(nums[end])
		for currSum*int64(end-start+1) >= k {
			currSum -= int64(nums[start])
			start++
		}

		res += int64(end - start + 1)
	}

	return res
}
