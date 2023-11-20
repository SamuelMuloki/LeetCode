package solutions

func CountPalindromicSubsequence(s string) int {
	first, last := [26]int{}, [26]int{}
	for i := 0; i < 26; i++ {
		first[i], last[i] = -1, -1
	}

	for i := 0; i < len(s); i++ {
		curr := s[i] - 'a'
		if first[curr] == -1 {
			first[curr] = i
		}

		last[curr] = i
	}

	ans := 0
	for i := 0; i < 26; i++ {
		if first[i] == -1 {
			continue
		}

		between := make(map[byte]bool)
		for j := first[i] + 1; j < last[i]; j++ {
			between[s[j]] = true
		}

		ans += len(between)
	}

	return ans
}
