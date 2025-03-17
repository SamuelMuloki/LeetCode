package solutions

func MaximumLength(s string) int {
	count := make(map[string]int)
	for start := 0; start < len(s); start++ {
		curr := []byte{}
		for end := start; end < len(s); end++ {
			if len(curr) == 0 || curr[len(curr)-1] == s[end] {
				curr = append(curr, s[end])
				count[string(curr)]++
			} else {
				break
			}
		}
	}

	ans := 0
	for val, num := range count {
		str := val
		if num >= 3 && len(str) > ans {
			ans = len(str)
		}
	}
	if ans == 0 {
		return -1
	}

	return ans
}
