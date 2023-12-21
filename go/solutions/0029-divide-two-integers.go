package solutions

import "math"

func Divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	ans := 0
	dvd, dvs := abs(dividend), abs(divisor)
	for dvs <= dvd {
		temp, mul := dvs, 1
		for dvd >= temp<<1 {
			temp <<= 1
			mul <<= 1
		}
		dvd -= temp
		ans += mul
	}

	if (dividend < 0) != (divisor < 0) {
		return -ans
	}

	return ans
}
