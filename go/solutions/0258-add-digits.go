package solutions

func AddDigits(num int) int {
	for (num / 10) != 0 {
		num = (num / 10) + (num % 10)
	}

	return num
}
