package solutions

func MinSteps(s string, t string) int {
	counts := [26]int{}
	for i := range s {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	ans := 0
	for i := 0; i < 26; i++ {
		if counts[i] > 0 {
			ans += counts[i]
		}
	}

	return ans
}
