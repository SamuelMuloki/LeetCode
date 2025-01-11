package solutions

func WordSubsets(words1 []string, words2 []string) []string {
	count := [26]int{}
	for _, word := range words2 {
		count2 := [26]int{}
		for _, ch := range word {
			count2[ch-'a']++
			count[ch-'a'] = max(count[ch-'a'], count2[ch-'a'])
		}
	}

	ans := []string{}
	for _, word := range words1 {
		count2 := [26]int{}
		for _, ch := range word {
			count2[ch-'a']++
		}
		for i := 0; i <= 26; i++ {
			if i == 26 {
				ans = append(ans, word)
				break
			}

			if count[i] > count2[i] {
				break
			}
		}
	}

	return ans
}
