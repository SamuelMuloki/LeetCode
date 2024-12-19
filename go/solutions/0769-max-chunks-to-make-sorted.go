package solutions

func MaxChunksToSorted(arr []int) int {
	n := len(arr)
	prefixMax := make([]int, n)
	copy(prefixMax, arr)
	suffixMin := make([]int, n)
	copy(suffixMin, arr)

	for i := 1; i < n; i++ {
		prefixMax[i] = max(prefixMax[i-1], prefixMax[i])
	}

	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = min(suffixMin[i+1], suffixMin[i])
	}

	chunks := 0
	for i := 0; i < n; i++ {
		if i == 0 || suffixMin[i] > prefixMax[i-1] {
			chunks++
		}
	}

	return chunks
}
