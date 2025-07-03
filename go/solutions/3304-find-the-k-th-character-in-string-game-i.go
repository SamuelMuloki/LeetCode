package solutions

func KthCharacter(k int) byte {
	var dfs func(str string) byte
	dfs = func(str string) byte {
		if len(str) >= k {
			return str[k-1]
		}

		s := []byte{}
		for i := 0; i < len(str); i++ {
			if str[i]+1 > 'z' {
				s = append(s, 'a')
			} else {
				s = append(s, str[i]+1)
			}
		}

		return dfs(str + string(s))
	}

	return dfs("a")
}
