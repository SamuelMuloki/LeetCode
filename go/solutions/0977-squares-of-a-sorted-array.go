package solutions

func SortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	start, end := 0, n-1
	for k := n - 1; start <= end; k-- {
		startProd := nums[start] * nums[start]
		endProd := nums[end] * nums[end]
		if startProd > endProd {
			ans[k] = startProd
			start++
		} else {
			ans[k] = endProd
			end--
		}
	}

	return ans
}
