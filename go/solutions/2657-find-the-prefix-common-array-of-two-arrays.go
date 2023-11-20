package solutions

func FindThePrefixCommonArray(A []int, B []int) []int {
	setA, setB := make(map[int]bool), make(map[int]bool)
	ans := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		setA[A[i]] = true
		setB[B[i]] = true
		for j := 0; j <= i; j++ {
			if setB[A[j]] {
				ans[i]++
			}
		}
	}

	return ans
}
