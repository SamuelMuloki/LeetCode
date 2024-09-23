package solutions

func LexicalOrder(n int) []int {
	result := []int{}

	var dfs func(curr, n int, result *[]int)
	dfs = func(curr, n int, result *[]int) {
		if curr > n {
			return
		}

		*result = append(*result, curr)

		for i := 0; i <= 9; i++ {
			next := curr*10 + i
			if next > n {
				return
			}
			dfs(next, n, result)
		}
	}

	for i := 1; i <= 9; i++ {
		dfs(i, n, &result)
	}

	return result
}
