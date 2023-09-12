package solutions

func MaxSum(nums []int, m int, k int) int64 {
	count := make(map[int]int, 0)
	maxSum, sum := 0, 0

	for i := 0; i < k; i++ {
		sum += nums[i]
		count[nums[i]]++
	}

	if len(count) >= m {
		maxSum = sum
	}

	for j := k; j < len(nums); j++ {
		sum += nums[j] - nums[j-k]
		count[nums[j]]++
		count[nums[j-k]]--

		if count[nums[j-k]] == 0 {
			delete(count, nums[j-k])
		}

		if len(count) >= m && sum > maxSum {
			maxSum = sum
		}
	}

	return int64(maxSum)
}
