package solutions

func MaximumEnergy(energy []int, k int) int {
	dp := make([]int, k)
	n := len(energy)
	res := energy[n-1]
	for i := 0; i < n; i++ {
		idx := i % k
		dp[idx] += energy[n-1-i]
		res = max(res, dp[idx])
	}
	return res
}
