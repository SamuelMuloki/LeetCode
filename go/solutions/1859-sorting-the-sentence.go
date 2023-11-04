package solutions

import "strings"

func SortSentence(s string) string {
	splitS := strings.Split(s, " ")
	arr := make([]string, len(splitS))
	for i := range splitS {
		idx := splitS[i][len(splitS[i])-1] - '0'
		arr[idx-1] = splitS[i][:len(splitS[i])-1]
	}

	return strings.Join(arr, " ")
}
