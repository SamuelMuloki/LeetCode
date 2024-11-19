package solutions

import (
	"math"
	"sort"
)

func MinimumMountainRemovals(nums []int) int {
	n := len(nums)
	lis := lisLength(nums)

	reversed := make([]int, n)
	copy(reversed, nums)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	lds := lisLength(reversed)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		lds[i], lds[j] = lds[j], lds[i]
	}

	removals := math.MaxInt32
	for i := 0; i < n; i++ {
		if lis[i] > 1 && lds[i] > 1 {
			removals = min(removals, n+1-lis[i]-lds[i])
		}
	}

	return removals
}

func lisLength(v []int) []int {
	lis := []int{v[0]}
	lisLen := make([]int, len(v))
	for i := range lisLen {
		lisLen[i] = 1
	}

	for i := 1; i < len(v); i++ {
		if v[i] > lis[len(lis)-1] {
			lis = append(lis, v[i])
		} else {
			index := sort.SearchInts(lis, v[i])
			lis[index] = v[i]
		}
		lisLen[i] = len(lis)
	}
	return lisLen
}
