package solutions

import "sort"

func SpecialArray(nums []int) int {
	sort.Ints(nums)
	N := len(nums)
	for i := 1; i <= N; i++ {
		k := sort.Search(N, func(k int) bool { return nums[k] >= i })
		if N-k == i {
			return i
		}
	}

	return -1
}
