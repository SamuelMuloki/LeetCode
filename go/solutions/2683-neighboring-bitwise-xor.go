package solutions

func DoesValidArrayExist(derived []int) bool {
	XOR := 0
	for _, num := range derived {
		XOR ^= num
	}

	return XOR == 0
}
