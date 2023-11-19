package solutions

func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	str1 := ""
	for i := range word1 {
		str1 += word1[i]
	}

	str2 := ""
	for i := range word2 {
		str2 += word2[i]
	}

	return str1 == str2
}
