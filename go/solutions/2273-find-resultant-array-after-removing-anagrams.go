package solutions

func RemoveAnagrams(words []string) []string {
	var isAnagram = func(w1, w2 string) bool {
		if len(w1) != len(w2) {
			return false
		}

		cnt1 := make(map[rune]int)
		cnt2 := make(map[rune]int)

		for _, ch := range w1 {
			cnt1[ch]++
		}

		for _, ch := range w2 {
			cnt2[ch]++
		}

		for k, c := range cnt1 {
			if cnt2[k] != c {
				return false
			}
		}

		return true
	}

	res := []string{words[0]}
	for i := 1; i < len(words); i++ {
		if isAnagram(words[i], res[len(res)-1]) {
			continue
		}

		res = append(res, words[i])
	}

	return res
}
