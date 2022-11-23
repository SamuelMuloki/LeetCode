package leetcode

/*
Qn -> Find the Sum of k Consecutive elements in array
Window sliding technique
*/
func MaxSum(arr []int, k int) int {
	max_sum := 0

	for i := 0; i < k; i++ {
		max_sum += arr[i]
	}

	window_sum := max_sum
	for j := k; j < len(arr); j++ {
		window_sum += arr[j] - arr[j-k]
		max_sum = findMax(window_sum, max_sum)
	}

	return max_sum
}
