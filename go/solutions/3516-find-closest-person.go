package solutions

func FindClosest(x int, y int, z int) int {
	if abs(x-z) > abs(y-z) {
		return 2
	} else if abs(x-z) < abs(y-z) {
		return 1
	}

	return 0
}
