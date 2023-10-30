package solutions

import (
	"math/bits"
	"sort"
)

func SortByBits(arr []int) []int {
	sort.SliceStable(arr, func(i, j int) bool {
		countI, countJ := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		return (countI < countJ) || (countI == countJ && arr[i] < arr[j])
	})

	return arr
}
