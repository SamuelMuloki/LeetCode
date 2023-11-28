package solutions

func SumOddLengthSubarrays(arr []int) int {
	n, ans := len(arr), 0
	for i := 0; i < n; i++ {
		left, right := i, n-i-1
		ans += arr[i] * (left/2 + 1) * (right/2 + 1)
		ans += arr[i] * ((left + 1) / 2) * ((right + 1) / 2)
	}

	return ans
}
