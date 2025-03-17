package solutions

func FindChampionII(n int, edges [][]int) int {
	count := make([]int, n)
	for _, edge := range edges {
		count[edge[1]]++
	}

	champ, champCount := -1, 0
	for i := 0; i < n; i++ {
		if count[i] == 0 {
			champCount++
			champ = i
		}
	}

	if champCount > 1 {
		return -1
	}

	return champ
}
