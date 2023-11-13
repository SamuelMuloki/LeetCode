package solutions

func MinSteps2(s string, t string) int {
	counts := [26]int{}
	for i := range s {
		counts[s[i]-'a']++
	}

	for i := range t {
		counts[t[i]-'a']--
	}

	ans := 0
	for i := 0; i < 26; i++ {
		if counts[i] != 0 {
			ans += abs(counts[i])
		}
	}

	return ans
}
