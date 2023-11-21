package solutions

func CountBadPairs(nums []int) int64 {
	set := make(map[int]int)
	ans := 0
	for i := range nums {
		ans += i - set[nums[i]-i]
		set[nums[i]-i]++
	}

	return int64(ans)
}
