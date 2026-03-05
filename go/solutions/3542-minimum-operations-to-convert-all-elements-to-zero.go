package solutions

func MinOperations11(nums []int) int {
	res := 0
	stack := []int{}

	for _, num := range nums {
		for len(stack) > 0 && stack[len(stack)-1] > num {
			stack = stack[:len(stack)-1]
		}

		if num > 0 && (len(stack) == 0 || stack[len(stack)-1] < num) {
			res++
			stack = append(stack, num)
		}
	}

	return res
}
