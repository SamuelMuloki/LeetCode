package solutions

func ScoreOfParentheses(s string) int {
	ans, st := 0, []int{}
	for i := range s {
		if s[i] == '(' {
			st = append(st, ans)
			ans = 0
		} else {
			ans = st[len(st)-1] + max(2*ans, 1)
			st = st[:len(st)-1]
		}
	}

	return ans
}
