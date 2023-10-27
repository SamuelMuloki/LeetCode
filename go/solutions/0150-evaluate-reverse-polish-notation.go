package solutions

import "strconv"

func EvalRPN(tokens []string) int {
	st := []int{}
	ops := map[byte]bool{'+': true, '-': true, '*': true, '/': true}
	for i := 0; i < len(tokens); i++ {
		if len(tokens[i][:]) == 1 && ops[tokens[i][0]] && len(st) > 1 {
			last, second := st[len(st)-1], st[len(st)-2]
			st = st[:len(st)-2]

			var res int
			if tokens[i][0] == '+' {
				res = last + second
			} else if tokens[i][0] == '-' {
				res = second - last
			} else if tokens[i][0] == '*' {
				res = last * second
			} else {
				res = second / last
			}

			st = append(st, res)
		} else {
			num, _ := strconv.Atoi(tokens[i])
			st = append(st, num)
		}
	}

	return st[0]
}
