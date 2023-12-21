package solutions

import (
	"sort"
	"strings"
)

func ArrangeWords(text string) string {
	words := strings.Fields(strings.ToLower(text))
	sort.SliceStable(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	str := strings.Join(words, " ")

	return strings.ToUpper(string(str[0])) + str[1:]
}
