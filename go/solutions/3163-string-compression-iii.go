package solutions

import "strconv"

func CompressedString(word string) string {
	count, comp := 0, []byte{}
	for i := 1; i <= len(word); i++ {
		count = 1
		for i < len(word) && word[i-1] == word[i] && count%9 != 0 {
			count++
			i++
		}

		comp = append(comp, []byte(strconv.Itoa(count))...)
		comp = append(comp, word[i-1])
	}

	return string(comp)
}
