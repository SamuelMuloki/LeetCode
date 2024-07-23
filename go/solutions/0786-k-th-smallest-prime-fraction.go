package solutions

func KthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	left, right := 0.0, 1.0
	ans := make([]int, 2)

	for left < right {
		mid := (left + right) / 2
		count := 0
		maxFraction := [2]int{0, 1}

		for i, j := 0, 1; i < n-1; i++ {
			for j < n && float64(arr[i])/float64(arr[j]) > mid {
				j++
			}
			count += n - j
			if j < n && float64(arr[i])/float64(arr[j]) > float64(maxFraction[0])/float64(maxFraction[1]) {
				maxFraction = [2]int{arr[i], arr[j]}
			}
		}

		if count == k {
			return maxFraction[:]
		} else if count < k {
			left = mid
		} else {
			right = mid
		}
	}

	return ans
}
