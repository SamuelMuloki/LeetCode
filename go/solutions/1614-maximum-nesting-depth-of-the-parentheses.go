package solutions

func MaxDepth2(s string) int {
	ans, curr := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			curr++
			ans = max(ans, curr)
		} else if s[i] == ')' {
			curr--
		}
	}

	return ans
}
