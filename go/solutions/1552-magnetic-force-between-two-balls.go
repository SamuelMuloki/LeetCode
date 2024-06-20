package solutions

import (
	"sort"
)

func MaxDistance(position []int, m int) int {
	n := len(position)
	var canPlaceBalls = func(x int) bool {
		prevBallPos := position[0]
		ballsPlaced := 1

		for i := 1; i < n && ballsPlaced < m; i++ {
			currPos := position[i]
			if currPos-prevBallPos >= x {
				ballsPlaced += 1
				prevBallPos = currPos
			}
		}

		return ballsPlaced == m
	}

	ans := 0
	sort.Ints(position)

	low, high := 1, position[n-1]/(m-1)
	for low <= high {
		mid := low + (high-low)/2
		if canPlaceBalls(mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}
