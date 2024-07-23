package solutions

import "math"

func IsNStraightHand(hand []int, groupSize int) bool {
	N := len(hand)
	if N%groupSize != 0 {
		return false
	}

	count := make(map[int]int)
	for i := range hand {
		count[hand[i]]++
	}

	for len(count) > 0 {
		currentCard := smallest(count)
		for i := 0; i < groupSize; i++ {
			if count[currentCard+i] == 0 {
				return false
			}
			count[currentCard+i]--
			if count[currentCard+i] < 1 {
				delete(count, currentCard+i)
			}
		}
	}

	return true
}

func smallest(set map[int]int) int {
	minV := math.MaxInt32
	for groupSize := range set {
		minV = min(minV, groupSize)
	}

	return minV
}
