package solutions

func IsZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	diff := make([]int, n+1)
	for _, query := range queries {
		diff[query[0]] -= 1
		diff[query[1]+1] += 1
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			diff[i] += diff[i-1]
		}
		nums[i] += diff[i]
		if nums[i] > 0 {
			return false
		}
	}

	return true
}
