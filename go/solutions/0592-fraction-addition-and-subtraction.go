package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func FractionAddition(expression string) string {
	n := len(expression)
	numerator, denominator := 0, 1

	for i := 0; i < n; {
		sign := 1
		if expression[i] == '-' {
			sign = -1
			i++
		} else if expression[i] == '+' {
			i++
		}

		start := i
		for i < n && expression[i] != '+' && expression[i] != '-' {
			i++
		}
		num, den := parseFraction(expression[start:i])

		commonDenominator := lcm(denominator, den)
		numerator = numerator*(commonDenominator/denominator) + sign*num*(commonDenominator/den)
		denominator = commonDenominator
	}

	commonDivisor := gcd(abs(numerator), denominator)
	numerator /= commonDivisor
	denominator /= commonDivisor

	return fmt.Sprintf("%d/%d", numerator, denominator)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func parseFraction(s string) (int, int) {
	parts := strings.Split(s, "/")
	numerator, _ := strconv.Atoi(parts[0])
	denominator, _ := strconv.Atoi(parts[1])
	return numerator, denominator
}
