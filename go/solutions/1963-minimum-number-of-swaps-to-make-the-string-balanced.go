package solutions

func MinSwaps(s string) int {
	open, close := 0, 0
	for i := range s {
		if open > 0 && s[i] == ']' {
			open -= 1
		} else if s[i] == '[' {
			open++
		} else if s[i] == ']' {
			close++
		}
	}

	return (open + 1) / 2
}
