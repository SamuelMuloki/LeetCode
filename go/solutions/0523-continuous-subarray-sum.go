package solutions

func CheckSubarraySum(nums []int, k int) bool {
	prefixMod := 0
	modSeen := make(map[int]int)
	modSeen[0] = -1

	for i := 0; i < len(nums); i++ {
		prefixMod = (prefixMod + nums[i]) % k
		if val, found := modSeen[prefixMod]; found {
			if i-val > 1 {
				return true
			}
		} else {
			modSeen[prefixMod] = i
		}
	}

	return false
}
