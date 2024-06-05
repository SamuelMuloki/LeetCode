package solutions

func CommonChars(words []string) []string {
	count := make([]map[byte]int, len(words))
	for i := range words {
		cnt := make(map[byte]int)
		for j := range words[i] {
			cnt[words[i][j]]++
		}
		count[i] = cnt
	}

	ans := make([]string, 0)
	for k, v := range count[0] {
		cnt := v
		for i := 1; i < len(count); i++ {
			if val, found := count[i][k]; found {
				cnt = min(cnt, val)
			} else {
				cnt = 0
				break
			}
		}

		for i := 0; i < cnt; i++ {
			ans = append(ans, string(k))
		}
	}

	return ans
}
