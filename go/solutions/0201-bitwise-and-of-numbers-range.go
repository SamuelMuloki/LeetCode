package solutions

func RangeBitwiseAnd(left int, right int) int {
	for right > left {
		right &= right - 1
	}

	return right
}
