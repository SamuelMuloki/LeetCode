package solutions

func GetMaximumXor(nums []int, maximumBit int) []int {
	pref, n := 0, len(nums)
	for i := 0; i < n; i++ {
		pref ^= nums[i]
	}

	ans := make([]int, n)
	mask := (1 << maximumBit) - 1
	for i := 0; i < n; i++ {
		ans[i] = pref ^ mask
		pref = pref ^ nums[n-1-i]
	}

	return ans
}
