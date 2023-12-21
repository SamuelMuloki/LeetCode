package solutions

import "math"

func RearrangeCharacters(s string, target string) int {
	set := [26]int{}
	for i := range target {
		set[target[i]-'a']++
	}

	count := [26]int{}
	for i := range s {
		count[s[i]-'a']++
	}

	ans := math.MaxInt
	for _, ch := range target {
		ans = min(ans, count[ch-'a']/set[ch-'a'])
	}

	return ans
}
