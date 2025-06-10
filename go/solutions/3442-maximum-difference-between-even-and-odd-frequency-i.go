package solutions

import "math"

func MaxDifference(s string) int {
	m := [26]int{}
	for _, ch := range s {
		m[ch-'a']++
	}

	j, k := 0, math.MaxInt
	for i := 0; i < 26; i++ {
		if m[i] == 0 {
			continue
		}

		if m[i]%2 == 1 {
			j = max(j, m[i])
		} else {
			k = min(k, m[i])
		}
	}

	return j - k
}
