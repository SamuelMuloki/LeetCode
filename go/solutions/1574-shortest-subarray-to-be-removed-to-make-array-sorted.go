package solutions

func FindLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	left, right := 0, n-1
	for left < n-1 && arr[left] <= arr[left+1] {
		left++
	}

	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}

	if left >= right {
		return 0
	}

	ans := min(n-left-1, right)
	for i, j := 0, right; i <= left && j < n; i++ {
		for j < n && arr[j] < arr[i] {
			j++
		}
		ans = min(ans, j-i-1)
	}

	return ans
}
