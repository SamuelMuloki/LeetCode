package solutions

func MaxFreqSum(s string) int {
	vowels := [26]int{}
	consonants := [26]int{}
	maxV, maxC := 0, 0
	for _, ch := range s {
		if isVowel(ch) {
			vowels[ch-'a']++
			maxV = max(maxV, vowels[ch-'a'])
		} else {
			consonants[ch-'a']++
			maxC = max(maxC, consonants[ch-'a'])
		}
	}

	return maxV + maxC
}
