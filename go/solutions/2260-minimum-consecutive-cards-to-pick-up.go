package solutions

import "math"

func MinimumCardPickup(cards []int) int {
	n := len(cards)
	lastIndex := [1000001]int{}
	for i := range lastIndex {
		lastIndex[i] = -1
	}

	ans := math.MaxInt
	for i := 0; i < n; i++ {
		if lastIndex[cards[i]] != -1 {
			ans = min(ans, i-lastIndex[cards[i]]+1)
		}
		lastIndex[cards[i]] = i
	}

	if ans == math.MaxInt {
		return -1
	}

	return ans
}
