package solutions

func EqualSubstring(s string, t string, maxCost int) int {
	start, count := 0, 0
	ans, curr := 0, 0
	for i := 0; i < len(s); i++ {
		count++
		curr += abs(int(s[i]-'a') - int(t[i]-'a'))
		if curr > maxCost {
			curr -= abs(int(s[start]-'a') - int(t[start]-'a'))
			start++
			count--
		}

		ans = max(ans, count)
	}

	return ans
}
