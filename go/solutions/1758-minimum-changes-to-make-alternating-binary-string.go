package solutions

func MinOperations2(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			if s[i] == '1' {
				ans++
			}
		} else {
			if s[i] == '0' {
				ans++
			}
		}
	}

	return min(ans, len(s)-ans)
}
