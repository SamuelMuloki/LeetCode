package solutions

func CountOfSubstrings(word string, k int) int64 {
	var isVowel = func(c rune) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	var atLeastK = func(word string, k int) int64 {
		var numValidSubstrings int64
		start, end := 0, 0
		vowelCount := make(map[rune]int)
		consonantCount := 0

		for end < len(word) {
			newLetter := rune(word[end])

			if isVowel(newLetter) {
				vowelCount[newLetter]++
			} else {
				consonantCount++
			}

			for len(vowelCount) == 5 && consonantCount >= k {
				numValidSubstrings += int64(len(word) - end)
				startLetter := rune(word[start])
				if isVowel(startLetter) {
					vowelCount[startLetter]--
					if vowelCount[startLetter] == 0 {
						delete(vowelCount, startLetter)
					}
				} else {
					consonantCount--
				}
				start++
			}

			end++
		}

		return numValidSubstrings
	}

	return atLeastK(word, k) - atLeastK(word, k+1)
}
