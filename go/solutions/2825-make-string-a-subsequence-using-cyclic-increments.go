package solutions

func CanMakeSubsequence(str1 string, str2 string) bool {
	j := len(str2) - 1
	for i := len(str1) - 1; i >= 0 && j >= 0; i-- {
		val := str1[i] + 1
		if val > 'z' {
			val = 'a'
		}

		if str1[i] == str2[j] || val == str2[j] {
			j--
			continue
		}
	}

	return j < 0
}
