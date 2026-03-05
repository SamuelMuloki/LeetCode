package solutions

func SmallestRepunitDivByK(k int) int {
	rem := 0
	for i := 1; i <= k; i++ {
		rem = (rem*10 + 1) % k
		if rem == 0 {
			return i
		}
	}

	return -1
}
