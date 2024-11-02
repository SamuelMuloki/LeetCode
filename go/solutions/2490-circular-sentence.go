package solutions

func IsCircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}

	for i := 0; i < len(sentence); i++ {
		if i > 2 && sentence[i-1] == ' ' && sentence[i-2] != sentence[i] {
			return false
		}
	}

	return true
}
