package solutions

func MinimumOneBitOperations(n int) int {
	if n == 0 {
		return 0
	}

	k, curr := 0, 1
	for curr*2 <= n {
		curr *= 2
		k++
	}

	return (1 << (k + 1)) - 1 - MinimumOneBitOperations(n^curr)
}
