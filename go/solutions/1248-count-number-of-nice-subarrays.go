package solutions

func NumberOfSubarrays(nums []int, k int) int {
	for i := 0; i < len(nums); i++ {
		nums[i] %= 2
	}

	prefixCount := make([]int, len(nums)+1)
	prefixCount[0] = 1
	ans, s := 0, 0
	for _, num := range nums {
		s += num
		if s >= k {
			ans += prefixCount[s-k]
		}
		prefixCount[s]++
	}

	return ans
}
