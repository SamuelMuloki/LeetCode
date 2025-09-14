package solutions

import "strings"

func Spellchecker(wordlist []string, queries []string) []string {
	m := make(map[string]bool)
	lower := make(map[string]string)
	for _, word := range wordlist {
		m[word] = true
		l := strings.ToLower(word)
		if _, ok := lower[l]; !ok {
			lower[l] = word
		}
	}

	res := make([]string, len(queries))
	for i, query := range queries {
		if m[query] {
			res[i] = query
			continue
		}

		l := strings.ToLower(query)
		if word, ok := lower[l]; ok {
			res[i] = word
			continue
		}

		w1 := l
		for _, word := range wordlist {
			w2 := strings.ToLower(word)
			if len(w1) != len(w2) {
				continue
			}

			found := true
			for i, w := range w2 {
				if rune(w1[i]) != w {
					if !(isVowel(rune(w1[i])) && isVowel(w)) {
						found = false
						break
					}
				}
			}

			if found {
				res[i] = word
				break
			}
		}
	}

	return res
}
