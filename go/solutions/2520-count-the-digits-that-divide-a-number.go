package solutions

func CountDigits(num int) int {
	ans, temp := 0, num
	for temp > 0 {
		digit := temp % 10
		if num%digit == 0 {
			ans++
		}
		temp /= 10
	}

	return ans
}
