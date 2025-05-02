package solutions

func PushDominoes(dominoes string) string {
	res := []rune(dominoes)
	s := "L" + dominoes + "R"
	left := 0
	for right := 1; right < len(s); right++ {
		if s[right] == '.' {
			continue
		}

		if s[left] == s[right] {
			for i := left + 1; i < right; i++ {
				res[i-1] = rune(s[left])
			}
		} else if s[left] == 'R' && s[right] == 'L' {
			for i, j := left+1, right-1; i < j; i, j = i+1, j-1 {
				res[i-1] = 'R'
				res[j-1] = 'L'
			}
		}
		left = right
	}

	return string(res)
}
