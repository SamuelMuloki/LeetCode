package solutions

import (
	"math"
)

func MinSubArrayLen(target int, nums []int) int {
	output := math.MaxInt
	l, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			output = min(output, i+1-l)
			sum -= nums[l]
			l++
		}
	}

	if output == math.MaxInt {
		return 0
	}

	return output
}
