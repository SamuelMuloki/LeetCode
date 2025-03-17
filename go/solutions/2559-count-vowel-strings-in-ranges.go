package solutions

func VowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	vowelMap := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	vowelCount := make([]int, n)
	isVowel := make([]bool, n)
	for i, word := range words {
		if i > 0 {
			vowelCount[i] = vowelCount[i-1]
		}

		if _, ok := vowelMap[word[0]]; !ok {
			continue
		}

		if _, ok := vowelMap[word[len(word)-1]]; !ok {
			continue
		}

		isVowel[i] = true
		vowelCount[i]++
	}

	ans := make([]int, len(queries))
	for i, query := range queries {
		count := vowelCount[query[1]] - vowelCount[query[0]]
		if isVowel[query[0]] {
			count++
		}
		ans[i] = count
	}

	return ans
}
