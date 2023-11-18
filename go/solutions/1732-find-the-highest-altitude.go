package solutions

func LargestAltitude(gain []int) int {
	sum := make([]int, len(gain))
	sum[0] = gain[0]
	maxA := max(sum[0], 0)
	for i := 1; i < len(gain); i++ {
		sum[i] = sum[i-1] + gain[i]
		maxA = max(maxA, sum[i])
	}

	return maxA
}
