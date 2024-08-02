package solutions

import "math"

func MinSwaps2(nums []int) int {
	minimumSwaps := math.MaxInt
	totalOnes := 0
	for _, num := range nums {
		totalOnes += num
	}

	onesCount := nums[0]
	end := 0
	for start := 0; start < len(nums); start++ {
		if start != 0 {
			onesCount -= nums[start-1]
		}

		for end-start+1 < totalOnes {
			end++
			onesCount += nums[end%len(nums)]
		}

		minimumSwaps = min(minimumSwaps, totalOnes-onesCount)
	}

	return minimumSwaps
}
