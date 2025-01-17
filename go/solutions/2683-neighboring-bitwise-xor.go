package solutions

func DoesValidArrayExist(derived []int) bool {
	sum := 0
	for _, num := range derived {
		sum += num
	}

	return sum%2 == 0
}
