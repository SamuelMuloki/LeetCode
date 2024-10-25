package solutions

func RemoveSubfolders(folder []string) []string {
	exists := make(map[string]bool)
	for i := range folder {
		exists[folder[i]] = true
	}

	ans := []string{}
	for i := range folder {
		curr, found := "", false
		for j := 0; j < len(folder[i]); j++ {
			curr += string(folder[i][j])
			if curr == folder[i] || j+1 < len(folder[i]) && folder[i][j+1] != '/' {
				continue
			}

			if _, ok := exists[curr]; ok {
				found = true
			}
		}

		if !found {
			ans = append(ans, curr)
		}
	}

	return ans
}
