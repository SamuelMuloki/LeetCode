package solutions

func SumSubarrayMins(arr []int) int {
	var ans int
	dp := make([]int, len(arr))
	stack := []int{}
	for i := range arr {
		var cur int
		for len(stack) != 0 && arr[i] < arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			cur = (i + 1) * arr[i]
		} else {
			cur = dp[stack[len(stack)-1]] + (i-stack[len(stack)-1])*arr[i]
		}

		stack = append(stack, i)
		dp[i] = cur
		ans = (ans + cur) % 1000000007
	}

	return ans
}
