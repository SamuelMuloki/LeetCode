package solutions

import "strconv"

func EvalRPN(tokens []string) int {
	st := []int{}
	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	for _, token := range tokens {
		if lambda, ok := ops[token]; ok {
			last, second := st[len(st)-1], st[len(st)-2]
			res := lambda(second, last)
			st = st[:len(st)-2]

			st = append(st, res)
		} else {
			num, _ := strconv.Atoi(token)
			st = append(st, num)
		}
	}

	return st[0]
}
