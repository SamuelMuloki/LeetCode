package solutions

import (
	"math"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func CharacterReplacement(s string, k int) int {
	set := make(map[rune]bool)
	runes := []rune(s)
	for i := range runes {
		set[runes[i]] = true
	}

	maxV := math.MinInt
	for ru := range set {
		count, currMax, pos := 0, 0, 0
		dp := make([]int, len(runes))
		inc := make([]int, len(runes))
		for i := 0; i < len(runes); i++ {
			inc[i]++
			if runes[i] != ru {
				count++
				dp[i]++
			}

			if count > k {
				currMax -= inc[pos]
				count -= dp[pos]
				pos++
			}

			currMax++
			maxV = utils.Max(maxV, currMax)
		}
	}

	return maxV
}
