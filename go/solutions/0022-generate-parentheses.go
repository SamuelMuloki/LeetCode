package solutions

func GenerateParenthesis(n int) []string {
	output := make([]string, 0)

	var backtrack func(l, r int, s string)
	backtrack = func(l, r int, s string) {
		if len(s) == 2*n {
			output = append(output, s)
		}

		if l < n {
			backtrack(l+1, r, s+"(")
		}

		if r < l {
			backtrack(l, r+1, s+")")
		}
	}

	backtrack(0, 0, "")

	return output
}
