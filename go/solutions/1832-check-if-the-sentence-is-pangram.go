package solutions

func CheckIfPangram(sentence string) bool {
	set := make(map[byte]int)
	for i := range sentence {
		set[sentence[i]]++
	}

	return len(set) == 26
}
