package solutions

func LetterCombinations(digits string) []string {
	output := make([]string, 0)

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

	var backtrack func(idx int, curr string)
	backtrack = func(idx int, curr string) {
		if len(curr) == len(digits) {
			output = append(output, curr)
			return
		}

		for i := 0; i < len(combs[digits[idx]]); i++ {
			backtrack(idx+1, curr+string(combs[digits[idx]][i]))
		}
	}

	if len(digits) > 0 {
		backtrack(0, "")
	}

	return output
}
