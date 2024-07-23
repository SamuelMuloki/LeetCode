package solutions

import "math/bits"

func MinOperations4(nums []int, k int) int {
	finalXor := 0
	for i := range nums {
		finalXor ^= nums[i]
	}

	return bits.OnesCount(uint(finalXor ^ k))
}
