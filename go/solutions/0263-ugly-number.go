package solutions

func IsUgly(n int) bool {
	if n < 1 {
		return false
	}

	for n%2 == 0 || n%3 == 0 || n%5 == 0 {
		if n%5 == 0 {
			n /= 5
		}

		if n%3 == 0 {
			n /= 3
		}

		if n%2 == 0 {
			n /= 2
		}
	}

	return n == 1
}
