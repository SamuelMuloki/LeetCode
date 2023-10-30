package solutions

func CountOdds(low int, high int) int {
	n := high - low + 1
	if n&1 == 0 || (low&1 == 0 && high&1 == 0) {
		return n / 2
	}

	return n/2 + 1
}
