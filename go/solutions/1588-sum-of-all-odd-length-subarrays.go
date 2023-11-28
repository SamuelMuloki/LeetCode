package solutions

func SumOddLengthSubarrays(arr []int) int {
	n, ans := len(arr), 0
	for left := 0; left < n; left++ {
		currSum := 0
		for right := left; right < n; right++ {
			currSum += arr[right]
			if (right-left+1)%2 == 1 {
				ans += currSum
			}
		}
	}

	return ans
}
