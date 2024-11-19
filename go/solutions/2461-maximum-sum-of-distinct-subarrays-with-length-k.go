package solutions

func MaximumSubarraySum(nums []int, k int) int64 {
	ans := 0
	start, sum := 0, 0
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
		sum += nums[i]
		for count[nums[i]] > 1 || len(count) > k {
			sum -= nums[start]
			count[nums[start]]--
			if count[nums[start]] == 0 {
				delete(count, nums[start])
			}
			start++
		}

		if len(count) == k {
			ans = max(ans, sum)
		}
	}

	return int64(ans)
}
