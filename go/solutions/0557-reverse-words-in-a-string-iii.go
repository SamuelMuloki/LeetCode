package solutions

func ReverseWordsIII(s string) string {
	runes := []rune(s)
	i, rLen := -1, len(runes)
	for idx := 0; idx <= rLen; idx++ {
		if idx == rLen || runes[idx] == ' ' {
			i += 1
			j := idx - 1
			for ; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			i = idx
		}
	}

	return string(runes)
}
