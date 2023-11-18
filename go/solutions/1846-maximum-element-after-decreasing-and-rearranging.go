package solutions

func MaximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)
	counts := make([]int, n+1)
	for i := range arr {
		counts[min(arr[i], n)]++
	}

	ans := 1
	for num := 2; num <= n; num++ {
		ans = min(ans+counts[num], num)
	}

	return ans
}
