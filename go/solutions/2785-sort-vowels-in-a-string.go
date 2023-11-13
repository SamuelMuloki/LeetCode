package solutions

import "sort"

func SortVowels(s string) string {
	temp := make([]rune, 0)
	runes := []rune(s)

	for i := range runes {
		if isVowel(runes[i]) {
			temp = append(temp, runes[i])
		}
	}

	sort.Slice(temp, func(i, j int) bool { return temp[i] < temp[j] })
	k := 0
	for i := 0; i < len(runes); i++ {
		if isVowel(runes[i]) {
			runes[i] = temp[k]
			k++
		}
	}

	return string(runes)
}
