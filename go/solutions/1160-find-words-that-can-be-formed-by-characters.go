package solutions

func CountCharacters(words []string, chars string) int {
	set := [26]int{}
	for i := range chars {
		set[chars[i]-'a']++
	}

	ans := 0
	for i := range words {
		found := true
		count := [26]int{}
		for j := range words[i] {
			count[words[i][j]-'a']++
			if count[words[i][j]-'a'] > set[words[i][j]-'a'] {
				found = false
				break
			}
		}

		if found {
			ans += len(words[i])
		}
	}

	return ans
}
