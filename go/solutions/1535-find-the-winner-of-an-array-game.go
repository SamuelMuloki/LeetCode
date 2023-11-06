package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func GetWinner(arr []int, k int) int {
	maxElement, n := arr[0], len(arr)
	for i := 1; i < n; i++ {
		maxElement = utils.Max(maxElement, arr[i])
	}

	curr, winStreak := arr[0], 0
	for i := 1; i < n; i++ {
		opponent := arr[i]

		if curr > opponent {
			winStreak++
		} else {
			curr = opponent
			winStreak = 1
		}

		if winStreak == k || curr == maxElement {
			return curr
		}
	}

	return -1
}
