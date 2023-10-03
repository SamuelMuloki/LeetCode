package solutions

import (
	"reflect"
	"sort"
)

func CloseStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	w1, w2 := make(map[rune]int), make(map[rune]int)
	r1, r2 := []rune(word1), []rune(word2)
	f1, f2 := make([]int, 0), make([]int, 0)
	for i := range r1 {
		w1[r1[i]]++
	}

	for i := range r2 {
		if w1[r2[i]] == 0 {
			return false
		}
		w2[r2[i]]++
	}

	for _, count1 := range w1 {
		f1 = append(f1, count1)
	}

	for _, count2 := range w2 {
		f2 = append(f2, count2)
	}

	sort.Ints(f1)
	sort.Ints(f2)

	return reflect.DeepEqual(f1, f2)
}
