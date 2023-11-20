package solutions

func GetLastMoment(n int, left []int, right []int) int {
	ans := 0
	for i := range left {
		ans = max(ans, left[i])
	}

	for i := range right {
		ans = max(ans, n-right[i])
	}

	return ans
}
