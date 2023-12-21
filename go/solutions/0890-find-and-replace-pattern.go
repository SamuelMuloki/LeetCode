package solutions

func FindAndReplacePattern(words []string, pattern string) []string {
	ans := []string{}
	for i := range words {
		wMap, pMap := [26]int{}, [26]int{}
		match := true
		for j := range words[i] {
			if wMap[words[i][j]-'a'] != pMap[pattern[j]-'a'] {
				match = false
				break
			}

			wMap[words[i][j]-'a'] = j + 1
			pMap[pattern[j]-'a'] = j + 1
		}

		if match {
			ans = append(ans, words[i])
		}
	}

	return ans
}
