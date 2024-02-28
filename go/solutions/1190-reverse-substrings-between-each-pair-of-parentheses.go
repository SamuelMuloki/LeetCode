package solutions

func ReverseParentheses(s string) string {
	st1, st2 := []int{}, []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st1 = append(st1, len(st2))
		} else if s[i] == ')' {
			last := st1[len(st1)-1]
			st2Last := st2[last:]
			arr := reverseBytes(st2Last)
			st2 = st2[:last]

			st1 = st1[:len(st1)-1]
			st2 = append(st2, arr...)
		} else {
			st2 = append(st2, s[i])
		}
	}

	return string(st2)
}

func reverseBytes(arr []byte) []byte {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
