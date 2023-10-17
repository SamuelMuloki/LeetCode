package solutions

func ContainsNearbyDuplicate(nums []int, k int) bool {
	set, indices := make(map[int]int), make(map[int]int)
	var abs = func(val int) int {
		if val < 0 {
			return -val
		}

		return val
	}

	for i := range nums {
		set[nums[i]]++
		if idx, ok := indices[nums[i]]; ok {
			if set[nums[i]] > 1 && abs(i-idx) <= k {
				return true
			}
		}

		indices[nums[i]] = i
	}

	return false
}
