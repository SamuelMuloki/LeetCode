package solutions

import "sort"

func Merge(nums1 []int, m int, nums2 []int, n int) {
	j := 0
	for i := m; i < len(nums1); i++ {
		nums1[i] = nums2[j]
		j++
	}

	sort.Ints(nums1)
}
