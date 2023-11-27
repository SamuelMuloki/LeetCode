package solutions

func FindWordsContaining(words []string, x byte) []int {
	ans := make([]int, 0)
	for i := range words {
		found := false
		for j := range words[i] {
			if words[i][j] == x {
				found = true
				break
			}
		}

		if found {
			ans = append(ans, i)
		}
	}

	return ans
}
