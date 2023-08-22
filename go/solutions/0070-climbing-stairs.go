package solutions

func ClimbStairs(n int) int {
	if n <= 3 {
		return n
	}

	a, b := 2, 3
	for ; n > 3; n-- {
		a, b = b, a+b
	}

	return b
}
