package solutions

func ConvertToTitle(columnNumber int) string {
	if columnNumber == 0 {
		return ""
	}

	columnNumber--
	charCode := 'A' + rune(columnNumber%26)

	return ConvertToTitle(columnNumber/26) + string(charCode)
}
