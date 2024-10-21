package solutions

func MaxUniqueSplit(s string) int {
	ans := 1
	count := map[string]bool{}

	var dfs func(i, t int)
	dfs = func(i, t int) {
		if i >= len(s) {
			ans = max(ans, t)
			return
		}

		for j := i + 1; j <= len(s); j++ {
			sub := s[i:j]
			if !count[sub] {
				count[sub] = true
				dfs(j, t+1)
				count[sub] = false
			}
		}
	}

	dfs(0, 0)
	return ans
}
