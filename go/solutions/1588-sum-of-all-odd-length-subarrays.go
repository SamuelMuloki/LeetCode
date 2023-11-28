package solutions

func SumOddLengthSubarrays(arr []int) int {
	n, ans := len(arr), 0
	for left := 0; left < n; left++ {
		for right := left; right < n; right++ {
			if (right-left+1)%2 == 1 {
				currSum := 0
				for i := left; i < right+1; i++ {
					currSum += arr[i]
				}
				ans += currSum
			}
		}
	}

	return ans
}
