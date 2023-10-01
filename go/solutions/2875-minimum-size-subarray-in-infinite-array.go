package solutions

import (
	"math"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func MinSizeSubarray(nums []int, target int) int {
	n, output := len(nums), math.MaxInt
	i, l, sum := 0, 0, 0
	prefixSum := 0
	var endIndex int

	for i := range nums {
		prefixSum += nums[i]
	}

	endIndex = ((target / prefixSum) + 2) * n
	for i < endIndex {
		sum += nums[i%n]
		for sum > target {
			sum -= nums[l%n]
			l++
		}

		if sum == target {
			output = utils.Min(output, i-l+1)
		}
		i++
	}

	if output == math.MaxInt {
		return -1
	}

	return output
}
