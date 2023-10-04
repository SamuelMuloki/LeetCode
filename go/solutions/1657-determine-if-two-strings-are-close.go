package solutions

import (
	"reflect"
	"sort"
)

func CloseStrings(word1 string, word2 string) bool {
	w1, w2, w3, w4 := [26]int{}, [26]int{}, [26]int{}, [26]int{}
	for i := range word1 {
		w1[int(word1[i]-'a')]++
		w3[int(word1[i]-'a')] = 1
	}

	for i := range word2 {
		w2[int(word2[i]-'a')]++
		w4[int(word2[i]-'a')] = 1
	}

	sort.Ints(w1[:])
	sort.Ints(w2[:])

	return reflect.DeepEqual(w1, w2) && reflect.DeepEqual(w3, w4)
}
