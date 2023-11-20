package solutions

func FindThePrefixCommonArray(A []int, B []int) []int {
	n, count := len(A), 0
	freq := make([]int, n+1)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		freq[A[i]]++
		if freq[A[i]] == 2 {
			count++
		}

		freq[B[i]]++
		if freq[B[i]] == 2 {
			count++
		}

		ans[i] = count
	}

	return ans
}
