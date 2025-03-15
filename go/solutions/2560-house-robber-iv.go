package solutions

import "math"

func MinCapability(nums []int, k int) int {
	var check = func(mid int) bool {
		possibleThefts := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				possibleThefts++
				i++
			}
		}

		return possibleThefts >= k
	}

	maxV := math.MinInt
	for i := 0; i < len(nums); i++ {
		maxV = max(maxV, nums[i])
	}

	l, r := 0, maxV
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
