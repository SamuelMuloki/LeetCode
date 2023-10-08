package solutions

func DifferenceOfSums(n int, m int) int {
	divSum, nonDivSum := 0, 0
	for i := 1; i <= n; i++ {
		if i%m == 0 {
			divSum += i
		} else {
			nonDivSum += i
		}
	}

	return nonDivSum - divSum
}
