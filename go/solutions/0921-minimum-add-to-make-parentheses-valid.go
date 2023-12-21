package solutions

func MinAddToMakeValid(s string) int {
	open, close := 0, 0
	for i := range s {
		if s[i] == '(' {
			open++
		} else if open > 0 && s[i] == ')' {
			open--
		} else {
			close++
		}
	}

	return open + close
}
