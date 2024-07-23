package solutions

func NumSubarraysWithSum(nums []int, goal int) int {
	n, ans := len(nums), 0
	seen := make(map[int]int)
	seen[0] = 1

	currSum := 0
	for r := 0; r < n; r++ {
		currSum += nums[r]

		if _, ok := seen[currSum-goal]; ok {
			ans += seen[currSum-goal]
		}

		seen[currSum] += 1
	}

	return ans
}
