package solutions

func SubtractProductAndSum(n int) int {
	prod, sum := 1, 0
	for ; n > 0; n /= 10 {
		prod *= n % 10
		sum += n % 10
	}

	return prod - sum
}
