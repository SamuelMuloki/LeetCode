package solutions

func PivotInteger(n int) int {
	total := (n * (n + 1)) / 2
	for i := 1; i <= n; i++ {
		before := (i * (i + 1)) / 2

		if before == total-before+i {
			return i
		}
	}

	return -1
}
