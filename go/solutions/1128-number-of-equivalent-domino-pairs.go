package solutions

func NumEquivDominoPairs(dominoes [][]int) int {
	res := 0
	cnt := make(map[[2]int]int)
	for _, dm := range dominoes {
		curr := [2]int{min(dm[0], dm[1]), max(dm[0], dm[1])}
		res += cnt[curr]
		cnt[curr]++
	}

	return res
}
