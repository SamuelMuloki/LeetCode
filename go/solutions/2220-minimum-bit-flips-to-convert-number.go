package solutions

func MinBitFlips(start int, goal int) int {
	count := 0
	for n := start ^ goal; n != 0; n &= n - 1 {
		count++
	}

	return count
}
