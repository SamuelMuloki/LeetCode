package solutions

func NumTilePossibilities(tiles string) int {
	ans := 0
	occ := make([]int, 26)

	var backtrack func(occ []int)
	backtrack = func(occ []int) {
		for i := 0; i < 26; i++ {
			if occ[i] > 0 {
				occ[i]--
				ans++
				backtrack(occ)
				occ[i]++
			}
		}
	}

	for i := range tiles {
		occ[tiles[i]-'A']++
	}

	backtrack(occ)
	return ans
}
