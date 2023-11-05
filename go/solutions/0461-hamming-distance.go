package solutions

func HammingDistance(x int, y int) int {
	count := 0
	for i := x ^ y; i != 0; i >>= 1 {
		if i&1 == 1 {
			count++
		}
	}

	return count
}
