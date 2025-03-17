package solutions

func AddSpaces(s string, spaces []int) string {
	n := len(s) + len(spaces)
	ans := make([]rune, n)

	stringIndex, spaceIndex := 0, 0
	for i := range ans {
		if spaceIndex < len(spaces) && i == spaces[spaceIndex]+spaceIndex {
			ans[i] = ' '
			spaceIndex++
			continue
		}

		ans[i] = rune(s[stringIndex])
		stringIndex++
	}

	return string(ans)
}
