package solutions

import (
	"fmt"
)

func Multiply(num1 string, num2 string) string {
	n1, n2 := []byte(num1), []byte(num2)
	l1, l2 := len(num1), len(num2)
	res := make([]int, l1+l2)

	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			n1i := toMulInt(string(n1[i]))
			n2i := toMulInt(string(n2[j]))
			mul := n1i * n2i
			p1 := i + j
			p2 := p1 + 1
			sum := res[p2] + mul
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}

	i := 0
	for i < len(res) && res[i] == 0 {
		i++
	}

	str := ""
	for i < len(res) {
		str += fmt.Sprintf("%d", res[i])
		i++
	}

	if len(str) == 0 {
		return "0"
	}

	return str
}

func toMulInt(s string) int {
	var n int = 0
	for _, ch := range []byte(s) {
		ch -= '0'
		n = (n * 10) + int(ch)
	}

	return n
}
