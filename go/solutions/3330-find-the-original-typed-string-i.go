package solutions

func PossibleStringCount(word string) int {
	res := 1
	for i := 1; i < len(word); i++ {
		if word[i-1] == word[i] {
			res++
		}
	}

	return res
}
