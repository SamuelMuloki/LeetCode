package solutions

import "strings"

func Spellchecker(wordlist []string, queries []string) []string {
	m := make(map[string]bool)
	lower := make(map[string]string)
	e := make(map[string]string)
	var convert = func(word string) string {
		runes := []rune(word)
		for i, r := range runes {
			if isVowel(r) {
				runes[i] = '*'
			}
		}

		return string(runes)
	}

	for _, word := range wordlist {
		m[word] = true
		l := strings.ToLower(word)
		if _, ok := lower[l]; !ok {
			lower[l] = word
		}

		c := convert(l)
		if _, ok := e[c]; !ok {
			e[c] = word
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

		if w, ok := e[convert(l)]; ok {
			res[i] = w
			continue
		}
	}

	return res
}
