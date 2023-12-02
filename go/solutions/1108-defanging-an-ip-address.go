package solutions

func DefangIPaddr(address string) string {
	runes := []rune(address)
	ans := []rune{}
	for i := range runes {
		if runes[i] == '.' {
			ans = append(ans, '[')
			ans = append(ans, runes[i])
			ans = append(ans, ']')
		} else {
			ans = append(ans, runes[i])
		}
	}

	return string(ans)
}
