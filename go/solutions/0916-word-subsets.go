package solutions

func WordSubsets(words1 []string, words2 []string) []string {
	maxFrequencies := make([]int, 26)
	lettersNeeded := make(map[int]struct{})

	countFrequencies := func(word string) []int {
		frequencies := make([]int, 26)
		for _, c := range word {
			idx := c - 'a'
			frequencies[idx]++
		}
		return frequencies
	}

	for _, word := range words2 {
		wordFrequencies := countFrequencies(word)
		for i := 0; i < 26; i++ {
			if wordFrequencies[i] > maxFrequencies[i] {
				maxFrequencies[i] = wordFrequencies[i]
				lettersNeeded[i] = struct{}{}
			}
		}
	}

	lettersNeededSlice := make([]int, 0, len(lettersNeeded))
	for i := range lettersNeeded {
		lettersNeededSlice = append(lettersNeededSlice, i)
	}

	universalWords := []string{}
	for _, word := range words1 {
		wordFrequencies := countFrequencies(word)
		isUniversal := true
		for _, i := range lettersNeededSlice {
			if wordFrequencies[i] < maxFrequencies[i] {
				isUniversal = false
				break
			}
		}
		if isUniversal {
			universalWords = append(universalWords, word)
		}
	}

	return universalWords
}
