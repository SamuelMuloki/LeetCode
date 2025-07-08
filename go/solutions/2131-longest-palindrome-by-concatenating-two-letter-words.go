package solutions

func LongestPalindrome3(words []string) int {
	set := make(map[string]int)
	for _, word := range words {
		set[word]++
	}

	center := false
	res := 0
	for word, count := range set {
		rev := string([]byte{word[1], word[0]})
		if rev == word {
			res += (count / 2) * 4
			if count%2 == 1 {
				center = true
			}
		} else if word < rev {
			pairs := min(count, set[rev])
			res += pairs * 4
		}
	}

	if center {
		res += 2
	}

	return res
}
