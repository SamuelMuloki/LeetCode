package solutions

func LargestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if int(num[i]-'0')%2 == 1 {
			return string(num[:i+1])
		}
	}

	return ""
}
