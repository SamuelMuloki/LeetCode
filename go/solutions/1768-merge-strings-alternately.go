package solutions

func MergeAlternately(word1 string, word2 string) string {
	s := ""

	i, j := 0, 0
	for ; i < len(word1) && j < len(word2); i, j = i+1, j+1 {
		s += string(word1[i]) + string(word2[j])
	}

	s += word1[i:] + word2[j:]
	return s
}
