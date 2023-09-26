package utils

import (
	"fmt"
	"math"
)

func MaxSum(arr []int, k int) int {
	max_sum := 0

	for i := 0; i < k; i++ {
		max_sum += arr[i]
	}

	window_sum := max_sum
	for j := k; j < len(arr); j++ {
		window_sum += arr[j] - arr[j-k]
		max_sum = Max(window_sum, max_sum)
	}

	return max_sum
}

// Reverse a string

func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func Min(x int, y int) int {
	if x < y {
		return x
	}

	return y
}

func Max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

func PrintAllSubsets(n, minMask int) {
	fmt.Printf("%d", 0)

	for mask := 0; mask <= minMask; mask++ {
		for k := 0; k < n; k++ {
			if (mask & (1 << k)) != 0 {
				fmt.Printf("%d", k+1)
			}
		}
	}
	fmt.Println()
}

func MaximumCost(i, mask, n int, dp [3][1 << 3]int, cost [][]int) int {
	if i >= n {
		return 0 // base case
	}

	// Check if the state has already been calculated
	if dp[i][mask] != -1 {
		// ans for the sub problem has been calculated return it
		return dp[i][mask]
	}

	// ans for the state has not been calculated before
	// make recursion call and store in the dp array
	ans := math.MinInt
	for k := 0; k < n; k++ {
		// check if kth person has been assigned a task or not
		if (mask & (1 << k)) == 0 {
			// kth person has not been assigned a task
			// add cost and make the kth bit of mask = 1
			ans = Max(ans, cost[i][k]+MaximumCost(i+1, (mask|1<<k), n, dp, cost))
		}
	}

	// store result of subproblem in the dp array
	dp[i][mask] = ans

	return dp[i][mask]
}
