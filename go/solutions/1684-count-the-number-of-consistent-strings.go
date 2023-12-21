package solutions

func CountConsistentStrings(allowed string, words []string) int {
	set := [26]bool{}
	for i := range allowed {
		set[allowed[i]-'a'] = true
	}

	ans := 0
	for i := range words {
		found := true
		for j := range words[i] {
			if !set[words[i][j]-'a'] {
				found = false
				break
			}
		}

		if found {
			ans++
		}
	}

	return ans
}
