package solutions

func MissingRolls(rolls []int, mean int, n int) []int {
	sum := 0
	for _, roll := range rolls {
		sum += roll
	}

	remainingSum := mean*(n+len(rolls)) - sum
	if remainingSum > 6*n || remainingSum < n {
		return []int{}
	}

	distributeMean, mod := remainingSum/n, remainingSum%n
	nElements := make([]int, n)
	for i := range nElements {
		nElements[i] = distributeMean
	}

	for i := 0; i < mod; i++ {
		nElements[i]++
	}

	return nElements
}
