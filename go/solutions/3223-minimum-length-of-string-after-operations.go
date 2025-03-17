package solutions

func MinimumLength2(s string) int {
	count := [26]int{}
	ans := 0
	for _, ch := range s {
		count[ch-'a']++
	}

	for i := 0; i < 26; i++ {
		if count[i] == 0 {
			continue
		}

		if count[i]%2 == 0 {
			ans += 2
		} else {
			ans += 1
		}
	}

	return ans
}
