package solutions

func CountOperations(num1 int, num2 int) int {
	res := 0
	for num1 != 0 && num2 != 0 {
		res += num1 / num2
		num1, num2 = num2, num1%num2
	}

	return res
}
