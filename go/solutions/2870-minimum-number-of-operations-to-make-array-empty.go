package solutions

import "math"

func MinEmptyOperations(nums []int) int {
	set := make(map[int]int)
	for i := range nums {
		set[nums[i]]++
	}

	ans := 0
	for _, count := range set {
		if count == 1 {
			return -1
		}

		ans += int(math.Ceil(float64(count) / 3))
	}

	return ans
}
