package solutions

import "sort"

func FindLeastNumOfUniqueInts(arr []int, k int) int {
	set := make(map[int]int)
	for i := range arr {
		set[arr[i]]++
	}

	freq := []int{}
	for i := range set {
		freq = append(freq, set[i])
	}

	sort.Ints(freq)
	elemsRem := 0
	for i := 0; i < len(freq); i++ {
		elemsRem += freq[i]
		if elemsRem > k {
			return len(freq) - i
		}
	}

	return 0
}
