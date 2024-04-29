package solutions

func MinOperations4(nums []int, k int) int {
	finalXor := 0
	for i := range nums {
		finalXor ^= nums[i]
	}

	count := 0
	for k > 0 || finalXor > 0 {
		if k%2 != finalXor%2 {
			count++
		}

		k /= 2
		finalXor /= 2
	}

	return count
}
