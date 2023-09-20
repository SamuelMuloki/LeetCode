// https://leetcode.com/problems/subarray-sum-equals-k
package solutions

func SubarraySum(nums []int, k int) int {
	n, count := len(nums), 0
	seen := make(map[int]int)
	seen[0] = 1

	currSum := 0
	for r := 0; r < n; r++ {
		currSum += nums[r]

		if _, ok := seen[currSum-k]; ok {
			count += seen[currSum-k]
		}

		seen[currSum] += 1
	}

	return count
}
