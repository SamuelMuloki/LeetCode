package solutions

func GarbageCollection(garbage []string, travel []int) int {
	set := make(map[byte]int)
	for i := range garbage {
		for j := range garbage[i] {
			set[garbage[i][j]]++
		}
	}

	ans := 0
	for i := 0; i < len(garbage); i++ {
		ans += len(garbage[i])

		if i > 0 {
			ans += len(set) * travel[i-1]
		}

		for j := 0; j < len(garbage[i]); j++ {
			set[garbage[i][j]]--

			if set[garbage[i][j]] == 0 {
				delete(set, garbage[i][j])
			}
		}
	}

	return ans
}
