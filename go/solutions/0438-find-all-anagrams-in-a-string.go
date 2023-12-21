package solutions

func FindAnagrams(s string, p string) []int {
	counts := make([]int, 26)
	for i := range p {
		counts[p[i]-'a']++
	}

	ans, start := []int{}, 0
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']--
		if i-start+1 == len(p) {
			found := true
			for i := 0; i < 26; i++ {
				if counts[i] != 0 {
					found = false
				}
			}
			if found {
				ans = append(ans, start)
			}
			counts[s[start]-'a']++
			start++
		}
	}

	return ans
}
