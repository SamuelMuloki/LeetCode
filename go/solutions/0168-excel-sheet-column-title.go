package solutions

func ConvertToTitle(columnNumber int) string {
	var output string
	for columnNumber > 0 {
		columnNumber--
		charCode := 'A' + rune(columnNumber%26)
		output = string(charCode) + output
		columnNumber /= 26
	}

	return output
}
