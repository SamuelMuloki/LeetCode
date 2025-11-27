package solutions

import "math"

func MaxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	prefixSum := int64(0)
	res := int64(math.MinInt64)
	kSum := make([]int64, k)
	for i := 0; i < k; i++ {
		kSum[i] = math.MaxInt64 / 2
	}
	kSum[k-1] = 0
	for i := 0; i < n; i++ {
		prefixSum += int64(nums[i])
		if prefixSum-kSum[i%k] > res {
			res = prefixSum - kSum[i%k]
		}
		if prefixSum < kSum[i%k] {
			kSum[i%k] = prefixSum
		}
	}

	return res
}
