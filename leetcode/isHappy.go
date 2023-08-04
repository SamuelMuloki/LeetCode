package leetcode

func IsHappy(n int) bool {
	isHappy := false

	for {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}

		n = sum

		isHappy = n == 1
		if isHappy || n == 37 {
			break
		}
	}

	return isHappy
}
