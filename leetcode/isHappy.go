// https://leetcode.com/problems/happy-number
package leetcode

func IsHappy(n int) bool {
	seenNumbers := make(map[int]bool)
	for n > 1 && !seenNumbers[n] {
		seenNumbers[n] = true
		n = pdiFunction(n, 0)
	}

	return n == 1
}

func pdiFunction(n int, sum int) int {
	if n == 0 {
		return sum
	}
	sum += (n % 10) * (n % 10)
	n = n / 10

	return pdiFunction(n, sum)
}
