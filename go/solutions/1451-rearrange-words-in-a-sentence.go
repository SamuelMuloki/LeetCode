package solutions

import "sort"

func ArrangeWords(text string) string {
	words := make([][]rune, 0)
	runes := []rune(text)
	runes[0] += 32
	for i, lastNonSpace := 0, 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			words = append(words, runes[lastNonSpace:i])
			lastNonSpace = i + 1
		}

		if i+1 == len(runes) {
			words = append(words, runes[lastNonSpace:i+1])
		}
	}

	sort.SliceStable(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	words[0][0] -= 32

	ans := ""
	for i := range words {
		if i > 0 {
			ans += " "
		}
		ans += string(words[i])
	}

	return ans
}
