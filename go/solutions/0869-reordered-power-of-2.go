package solutions

func ReorderedPowerOf2(n int) bool {
	var getDigits = func(n int) [10]int {
		digits := [10]int{}

		for n > 0 {
			digits[n%10]++
			n /= 10
		}

		return digits
	}

	digits := getDigits(n)
	for i := 1; i <= 1_000_000_000; i *= 2 {
		if getDigits(i) == digits {
			return true
		}
	}

	return false
}
