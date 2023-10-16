package solutions

func Tribonacci(n int) int {
	dp := make(map[int]int, n)
	var trib func(num int) int
	trib = func(num int) int {
		if num == 0 {
			return 0
		}

		if num <= 2 {
			return 1
		}

		if val, found := dp[num]; found {
			return val
		}

		dp[num] = trib(num-3) + trib(num-2) + trib(num-1)
		return dp[num]
	}

	return trib(n)
}
