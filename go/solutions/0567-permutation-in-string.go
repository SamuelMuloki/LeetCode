package solutions

func CheckInclusion(s1 string, s2 string) bool {
	counts := make([]int, 26)
	for i := range s1 {
		counts[s1[i]-'a']++
	}

	start := 0
	for i := 0; i < len(s2); i++ {
		counts[s2[i]-'a']--
		if i-start+1 == len(s1) {
			found := true
			for i := 0; i < 26; i++ {
				if counts[i] != 0 {
					found = false
				}
			}
			if found {
				return true
			}
			counts[s2[start]-'a']++
			start++
		}
	}

	return false
}
