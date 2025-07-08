package solutions

func RobotWithString(s string) string {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}

	stack := []rune{}
	res := []rune{}
	minCharacter := 'a'
	for _, c := range s {
		stack = append(stack, c)
		cnt[c-'a']--
		for minCharacter != 'z' && cnt[minCharacter-'a'] == 0 {
			minCharacter++
		}
		for len(stack) > 0 && stack[len(stack)-1] <= minCharacter {
			res = append(res, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	return string(res)
}
