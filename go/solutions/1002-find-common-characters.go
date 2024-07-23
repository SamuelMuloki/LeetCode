package solutions

func CommonChars(words []string) []string {
	wordsSize := len(words)
	var currentCharacterCounts [26]int
	commonCharacterCounts := [26]int{}
	for _, ch := range words[0] {
		commonCharacterCounts[ch-'a']++
	}

	for i := 1; i < wordsSize; i++ {
		currentCharacterCounts = [26]int{}
		for _, ch := range words[i] {
			currentCharacterCounts[ch-'a']++
		}

		for letter := 0; letter < 26; letter++ {
			commonCharacterCounts[letter] = min(commonCharacterCounts[letter], currentCharacterCounts[letter])
		}
	}

	ans := make([]string, 0)
	for letter := 0; letter < 26; letter++ {
		for commonCount := 0; commonCount < commonCharacterCounts[letter]; commonCount++ {
			ans = append(ans, string(byte(letter+'a')))
		}
	}

	return ans
}
