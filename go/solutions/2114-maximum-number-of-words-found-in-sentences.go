package solutions

func MostWordsFound(sentences []string) int {
	ans := 0
	for i := range sentences {
		count := 1
		for j := range sentences[i] {
			if sentences[i][j] == ' ' {
				count++
			}
		}

		ans = max(ans, count)
	}

	return ans
}
