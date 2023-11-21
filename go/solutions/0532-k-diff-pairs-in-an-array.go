package solutions

func FindPairs(nums []int, k int) int {
	ans := 0
	set := make(map[int]int)
	for i := range nums {
		set[nums[i]]++
	}

	for i := range set {
		val, ok := set[i+k]
		if k == 0 {
			if val > 1 {
				ans++
			}
		} else if ok {
			ans++
		}
	}

	return ans
}
