package solutions

import "math"

func MinimumSubarrayLength(nums []int, k int) int {
	minLength := math.MaxInt
	windowStart, windowEnd := 0, 0
	bitCounts := [32]int{}

	var updateBitCounts = func(number, delta int) {
		for bitPosition := 0; bitPosition < 32; bitPosition++ {
			if ((number >> bitPosition) & 1) != 0 {
				bitCounts[bitPosition] += delta
			}
		}
	}

	var converBitCountsToNumber = func() int {
		result := 0
		for bitPosition := 0; bitPosition < 32; bitPosition++ {
			if bitCounts[bitPosition] != 0 {
				result |= 1 << bitPosition
			}
		}

		return result
	}

	for windowEnd < len(nums) {
		updateBitCounts(nums[windowEnd], 1)

		for windowStart <= windowEnd && converBitCountsToNumber() >= k {
			minLength = min(minLength, windowEnd-windowStart+1)
			updateBitCounts(nums[windowStart], -1)
			windowStart++
		}
		windowEnd++
	}

	if minLength == math.MaxInt {
		return -1
	}

	return minLength
}
