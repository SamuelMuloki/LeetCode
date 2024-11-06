package solutions

import (
	"math/bits"
)

func CanSortArray(nums []int) bool {
	n := len(nums)
	arr := make([]int, n)
	copy(arr, nums)

	for i := 0; i < n-1; i++ {
		if arr[i] <= arr[i+1] {
			continue
		} else {
			if bits.OnesCount(uint(arr[i])) == bits.OnesCount(uint(arr[i+1])) {
				swapArr(arr, i, i+1)
			} else {
				return false
			}
		}
	}

	for i := n - 1; i > 0; i-- {
		if arr[i] >= arr[i-1] {
			continue
		} else {
			if bits.OnesCount(uint(arr[i])) == bits.OnesCount(uint(arr[i-1])) {
				swapArr(arr, i, i-1)
			} else {
				return false
			}
		}
	}

	return true
}
