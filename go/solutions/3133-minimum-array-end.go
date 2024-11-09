package solutions

func MinEnd(n int, x int) int64 {
	result := x
	n -= 1
	for mask := 1; n > 0; mask <<= 1 {
		if (mask & x) == 0 {
			result |= (n & 1) * mask
			n >>= 1
		}
	}

	return int64(result)
}
