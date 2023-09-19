package solutions

import "sort"

func CombinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	var backtrack func(currSum int, dp map[int]int) int
	backtrack = func(currSum int, dp map[int]int) int {
		if currSum == 0 {
			return 1
		}

		if currSum < 0 {
			return 0
		}

		if val, found := dp[currSum]; found {
			return val
		}

		take := 0
		for i := 0; i < len(nums); i++ {
			if currSum-nums[i] < 0 {
				break
			}
			take += backtrack(currSum-nums[i], dp)
		}

		dp[currSum] = take

		return dp[currSum]
	}

	return backtrack(target, make(map[int]int))
}
