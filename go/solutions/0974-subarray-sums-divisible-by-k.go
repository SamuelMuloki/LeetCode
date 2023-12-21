package solutions

func SubarraysDivByK(nums []int, k int) int {
	ans, prefixMod := 0, 0

	modGroups := make([]int, k)
	modGroups[0] = 1

	for i := range nums {
		prefixMod = (prefixMod + nums[i]%k + k) % k
		ans += modGroups[prefixMod]
		modGroups[prefixMod]++
	}

	return ans
}
