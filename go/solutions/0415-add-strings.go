package solutions

import "fmt"

func AddStrings(num1 string, num2 string) string {
	ret, carry := "", 0

	for x, y := len(num1)-1, len(num2)-1; x >= 0 || y >= 0; x, y = x-1, y-1 {
		i, j := toInt(x, num1), toInt(y, num2)
		sum := i + j + carry
		ret = fmt.Sprintf("%d%s", sum%10, ret)

		if sum <= 9 {
			carry = 0
		} else {
			carry = 1
		}
	}

	if carry == 1 {
		ret = fmt.Sprintf("1%s", ret)
	}

	return ret
}

func toInt(i int, s string) int {
	if i < 0 {
		return 0
	}
	return int(s[i] - '0')
}
