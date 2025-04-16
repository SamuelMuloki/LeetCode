package solutions

func CountGood(nums []int, k int) int64 {
	cnt := make(map[int]int)
	start := 0
	n := len(nums)
	var res int
	pairs := 0
	for end := 0; end < n; end++ {
		pairs += cnt[nums[end]]
		cnt[nums[end]]++

		for start < n && pairs >= k {
			res += n - end
			cnt[nums[start]]--
			pairs -= cnt[nums[start]]
			start++
		}
	}

	return int64(res)
}
