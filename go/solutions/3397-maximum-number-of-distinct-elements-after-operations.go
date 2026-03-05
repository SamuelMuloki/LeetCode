package solutions

import (
	"sort"
)

func MaxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	curr := -1 << 31
	res := 0
	for _, num := range nums {
		l, r := num-k, num+k
		candidate := max(curr+1, l)
		if candidate <= r {
			curr = candidate
			res++
		}
	}

	return res
}
