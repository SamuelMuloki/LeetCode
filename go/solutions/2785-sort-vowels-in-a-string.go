package solutions

func SortVowels(s string) string {
	count := make(map[rune]int)
	runes := []rune(s)

	for i := range runes {
		if isVowel(runes[i]) {
			count[runes[i]]++
		}
	}

	k, vowels := 0, "AEIOUaeiou"
	sortedVowels := []rune(vowels)
	for i := 0; i < len(runes); i++ {
		if !isVowel(runes[i]) {
			continue
		}

		for count[sortedVowels[k]] == 0 {
			k++
		}

		runes[i] = sortedVowels[k]
		count[runes[i]]--
	}

	return string(runes)
}
