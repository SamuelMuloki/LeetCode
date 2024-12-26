package solutions

import "math"

func FindTargetSumWays(nums []int, target int) int {
	totalSum := 0
	for i := range nums {
		totalSum += nums[i]
	}

	memo := make([][]int, len(nums))
	for i := range memo {
		memo[i] = make([]int, 2*totalSum+1)
		for j := range memo[i] {
			memo[i][j] = math.MaxInt
		}
	}

	var calculateWays func(currentIndex, currSum int) int
	calculateWays = func(currentIndex, currSum int) int {
		if currentIndex == len(nums) {
			if currSum == target {
				return 1
			} else {
				return 0
			}
		} else {

			if memo[currentIndex][currSum+totalSum] != math.MaxInt {
				return memo[currentIndex][currSum+totalSum]
			}

			add := calculateWays(
				currentIndex+1,
				currSum+nums[currentIndex],
			)

			subtract := calculateWays(
				currentIndex+1,
				currSum-nums[currentIndex],
			)

			memo[currentIndex][currSum+totalSum] = add + subtract
			return memo[currentIndex][currSum+totalSum]
		}
	}

	return calculateWays(0, 0)
}
