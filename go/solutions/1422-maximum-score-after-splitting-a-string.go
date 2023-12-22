package solutions

func MaxScore(s string) int {
	n := len(s)
	left, right := make([]int, n), make([]int, n)
	left[0] = int(s[0]-'0') ^ 1
	right[n-1] = int(s[n-1] - '0')
	ans := max(0, left[0]+right[n-1])
	for i := 1; i < n; i++ {
		left[i] = left[i-1]
		left[i] += int(s[i]-'0') ^ 1

		right[n-i-1] = right[n-i]
		right[n-i-1] += int(s[n-i-1] - '0')
	}

	for i := 1; i < n-1; i++ {
		ans = max(ans, left[i]+right[i])
	}

	return ans
}
