package solutions

func TrailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	return (n / 5) + TrailingZeroes(n/5)
}
