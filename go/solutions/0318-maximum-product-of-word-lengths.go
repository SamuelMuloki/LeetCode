package solutions

func MaxProduct2(words []string) int {
	ans, chars := 0, make([]uint32, len(words))
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			chars[i] |= 1 << (words[i][j] - 'a')
		}

		for j := 0; j < i; j++ {
			if chars[i]&chars[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}

	return ans
}
