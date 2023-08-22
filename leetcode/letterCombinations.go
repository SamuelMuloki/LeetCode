package leetcode

func LetterCombinations(digits string) []string {
	output := make([]string, 0)
	if len(digits) == 0 {
		return output
	}

	curr := ""
	var combs = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(curr) == len(digits) {
			output = append(output, curr)
			return
		}

		dig := digits[idx]
		str := combs[dig]
		for i := 0; i < len(str); i++ {
			curr += string(str[i])
			backtrack(idx + 1)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0)

	return output
}
