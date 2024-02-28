package solutions

func MinOperations3(boxes string) []int {
	n := len(boxes)
	leftSum, rightSum := make([]int, n), make([]int, n)
	leftSum[0], rightSum[n-1] = 0, 0
	lBalls, rBalls := 0, 0
	for i := 1; i < n; i++ {
		leftSum[i] = leftSum[i-1]
		lBalls += int(boxes[i-1] - '0')
		leftSum[i] += lBalls

		rightSum[n-i-1] = rightSum[n-i]
		rBalls += int(boxes[n-i] - '0')
		rightSum[n-i-1] += rBalls
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = leftSum[i] + rightSum[i]
	}

	return ans
}
