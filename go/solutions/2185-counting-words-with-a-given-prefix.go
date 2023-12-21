package solutions

func PrefixCount(words []string, pref string) int {
	ans := 0
	for _, word := range words {
		if len(pref) <= len(word) && word[:len(pref)] == pref {
			ans++
		}
	}

	return ans
}
