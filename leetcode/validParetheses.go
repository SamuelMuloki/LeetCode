package leetcode

func IsValid(s string) bool {
	runes := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	chars := []rune{}
	for _, char := range s {
		if _, ok := runes[char]; ok {
			chars = append(chars, char)
		} else {
			if len(chars) > 0 && runes[chars[len(chars)-1]] == char {
				chars = chars[:len(chars)-1]
			} else {
				return false
			}
		}
	}

	return len(chars) == 0
}
