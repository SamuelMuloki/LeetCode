package solutions

func FirstUniqChar(s string) int {
	set := make(map[rune]int)

	runes := []rune(s)
	for i := range runes {
		set[runes[i]]++
	}

	for i, ch := range runes {
		if set[ch] == 1 {
			return i
		}
	}

	return -1
}
