package solutions

func DoesValidArrayExist(derived []int) bool {
	n := len(derived)
	original := make([]int, n)
	for i := 0; i < n-1; i++ {
		original[i+1] = derived[i] ^ original[i]
	}

	checkForZero := derived[n-1] == (original[n-1] ^ original[0])

	original = make([]int, n)
	original[0] = 1
	for i := 0; i < n-1; i++ {
		original[i+1] = derived[i] ^ original[i]
	}

	checkForOne := derived[n-1] == (original[n-1] ^ original[0])

	return checkForZero || checkForOne
}
