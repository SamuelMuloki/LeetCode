package solutions

func SimilarPairs(words []string) int {
	ans := 0
	set := make(map[[26]byte]int)
	for i := range words {
		wSet := [26]byte{}
		for j := range words[i] {
			wSet[words[i][j]-'a'] = 1
		}

		ans += set[wSet]
		set[wSet]++
	}

	return ans
}
