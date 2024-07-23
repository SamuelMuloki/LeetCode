package solutions

func MaxScoreWords(words []string, letters []byte, score []int) int {
	wLen := len(words)
	freq := [26]int{}
	for _, c := range letters {
		freq[c-'a']++
	}

	var subsetScore = func(subsetLetters [26]int) int {
		totalScore := 0
		for c := 0; c < 26; c++ {
			totalScore += subsetLetters[c] * score[c]
			if subsetLetters[c] > freq[c] {
				return 0
			}
		}

		return totalScore
	}

	ans := 0
	for mask := 0; mask < (1 << wLen); mask++ {
		subsetLetters := [26]int{}
		for i := 0; i < wLen; i++ {
			if (mask & (1 << i)) > 0 {
				L := len(words[i])
				for j := 0; j < L; j++ {
					subsetLetters[words[i][j]-'a']++
				}
			}
		}
		ans = max(ans, subsetScore(subsetLetters))
	}

	return ans
}
