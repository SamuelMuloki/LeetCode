package solutions

func HammingDistance(x int, y int) int {
	count := 0
	// n & (n-1) converts the right most 1 to 0
	for n := x ^ y; n != 0; n &= n - 1 {
		count++
	}

	return count
}
