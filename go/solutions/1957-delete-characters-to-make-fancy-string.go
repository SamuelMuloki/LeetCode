package solutions

func MakeFancyString(s string) string {
	start, count := 0, 0
	ans := []byte{}
	for i := 0; i < len(s); i++ {
		if s[start] == s[i] {
			count++
		} else {
			count = 1
			start = i
		}

		if count < 3 {
			ans = append(ans, s[i])
		}
	}

	return string(ans)
}
