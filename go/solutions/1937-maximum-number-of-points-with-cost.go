package solutions

func MaxPoints(points [][]int) int64 {
	m, n := len(points), len(points[0])
	dp := make([]int64, n)

	for j := 0; j < n; j++ {
		dp[j] = int64(points[0][j])
	}

	for i := 1; i < m; i++ {
		leftMax := make([]int64, n)
		rightMax := make([]int64, n)
		newDp := make([]int64, n)

		leftMax[0] = dp[0]
		for j := 1; j < n; j++ {
			leftMax[j] = max(leftMax[j-1], dp[j]+int64(j))
		}

		rightMax[n-1] = dp[n-1] - int64(n-1)
		for j := n - 2; j >= 0; j-- {
			rightMax[j] = max(rightMax[j+1], dp[j]-int64(j))
		}

		for j := 0; j < n; j++ {
			newDp[j] = max(leftMax[j]-int64(j), rightMax[j]+int64(j)) + int64(points[i][j])
		}

		dp = newDp
	}

	result := dp[0]
	for j := 1; j < n; j++ {
		result = max(result, dp[j])
	}
	return result
}
