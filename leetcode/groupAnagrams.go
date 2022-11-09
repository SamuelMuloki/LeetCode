package leetcode

import "sort"

func GroupAnagrams(strs []string) [][]string {
	anagrams := [][]string{}
	evaluated := make(map[string][]string)

	for i := range strs {
		key := sortString(strs[i])
		evaluated[key] = append(evaluated[key], strs[i])
	}

	for idx := range evaluated {
		anagrams = append(anagrams, evaluated[idx])
	}

	return anagrams
}

func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(a, b int) bool {
		return runes[a] < runes[b]
	})
	return string(runes)
}
