package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func ConstrainedSubsetSum(nums []int, k int) int {
	var queue []int
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if len(queue) > 0 && i-queue[0] > k {
			queue = queue[1:]
		}

		dp[i] = nums[i]
		if len(queue) > 0 {
			dp[i] += dp[queue[0]]
		}

		for len(queue) > 0 && dp[queue[len(queue)-1]] < dp[i] {
			queue = queue[:len(queue)-1]
		}

		if dp[i] > 0 {
			queue = append(queue, i)
		}
	}

	return utils.MaxArr(dp)
}
