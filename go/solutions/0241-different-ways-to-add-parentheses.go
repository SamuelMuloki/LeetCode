package solutions

import "strconv"

func DiffWaysToCompute(expression string) []int {
	var res []int
	for i := 0; i < len(expression); i++ {
		oper := expression[i]
		if oper == '+' || oper == '-' || oper == '*' {
			s1 := DiffWaysToCompute(expression[:i])
			s2 := DiffWaysToCompute(expression[i+1:])
			for _, a := range s1 {
				for _, b := range s2 {
					if oper == '+' {
						res = append(res, a+b)
					} else if oper == '-' {
						res = append(res, a-b)
					} else if oper == '*' {
						res = append(res, a*b)
					}
				}
			}
		}
	}

	if len(res) == 0 {
		num, _ := strconv.Atoi(expression)
		res = append(res, num)
	}

	return res
}
