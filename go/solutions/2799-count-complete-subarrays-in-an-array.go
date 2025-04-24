package solutions

func CountCompleteSubarrays(nums []int) int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	start := 0
	curr := make(map[int]int)
	ans := 0
	n := len(nums)
	for end := 0; end < n; end++ {
		curr[nums[end]]++
		for len(curr) == len(cnt) {
			ans += n - end
			curr[nums[start]]--
			if curr[nums[start]] == 0 {
				delete(curr, nums[start])
			}
			start++
		}
	}

	return ans
}
