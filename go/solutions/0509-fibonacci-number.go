package solutions

func Fib(n int) int {
	if n <= 1 {
		return n
	}

	if n == 2 || n == 3 {
		return n - 1
	}

	prev, curr := 1, 2
	for k := 0; k < n-3; k++ {
		newFib := prev + curr
		prev = curr
		curr = newFib
	}

	return curr
}
