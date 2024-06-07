package solutions

import "strings"

func ReplaceWords(dictionary []string, sentence string) string {
	words := strings.Fields(sentence)
	for i := range words {
		val := ""
		for j := range dictionary {
			if strings.HasPrefix(words[i], dictionary[j]) {
				if val == "" || len(val) > len(dictionary[j]) {
					val = dictionary[j]
				}
			}
		}

		if val != "" {
			words[i] = val
		}
	}

	return strings.Join(words, " ")
}
