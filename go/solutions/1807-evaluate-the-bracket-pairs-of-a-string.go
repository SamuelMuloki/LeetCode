package solutions

func Evaluate(s string, knowledge [][]string) string {
	set := make(map[string][]rune)
	for i := range knowledge {
		set[knowledge[i][0]] = []rune(knowledge[i][1])
	}

	ans, start := []rune{}, -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			start = i
		}

		if s[i] == ')' {
			if val, ok := set[s[start+1:i]]; ok {
				ans = append(ans, val...)
			} else {
				ans = append(ans, '?')
			}

			start = -1
			continue
		}

		if start != -1 {
			continue
		}

		ans = append(ans, rune(s[i]))
	}

	return string(ans)
}
