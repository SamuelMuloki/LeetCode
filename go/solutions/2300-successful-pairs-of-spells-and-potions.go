package solutions

import "sort"

func SuccessfulPairs(spells []int, potions []int, success int64) []int {
	n := len(potions)
	m := len(spells)
	res := make([]int, m)

	sort.Ints(potions)
	for i := 0; i < m; i++ {
		spell := spells[i]
		minPotionStrength := (success + int64(spell) - 1) / int64(spell)

		idx := sort.Search(n, func(k int) bool { return int64(potions[k]) >= minPotionStrength })
		res[i] = n - idx
	}

	return res
}
