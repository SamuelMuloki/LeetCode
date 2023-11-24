package solutions

import "sort"

func MaxCoins(piles []int) int {
	sort.Ints(piles)
	ans := 0
	for i := len(piles) - 1; i >= len(piles)/3; i -= 2 {
		ans += piles[i-1]
	}

	return ans
}
