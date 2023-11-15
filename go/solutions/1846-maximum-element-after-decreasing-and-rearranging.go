package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func MaximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)
	counts := make([]int, n+1)
	for i := range arr {
		counts[utils.Min(arr[i], n)]++
	}

	ans := 1
	for num := 2; num <= n; num++ {
		ans = utils.Min(ans+counts[num], num)
	}

	return ans
}
