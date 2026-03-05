package solutions

func MaximumUniqueSubarray(nums []int) int {
	cnt := make(map[int]int)
	start := 0
	currSum := 0
	res := 0
	for end := 0; end < len(nums); end++ {
		cnt[nums[end]]++
		for start < end && cnt[nums[end]] > 1 {
			currSum -= nums[start]
			cnt[nums[start]]--
			start++
		}

		currSum += nums[end]
		res = max(res, currSum)
	}

	return res
}
