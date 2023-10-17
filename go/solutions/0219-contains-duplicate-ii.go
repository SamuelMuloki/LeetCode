package solutions

func ContainsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]int)
	for i := range nums {
		if idx, ok := set[nums[i]]; ok {
			if idx >= i-k {
				return true
			}
		}

		set[nums[i]] = i
	}

	return false
}
