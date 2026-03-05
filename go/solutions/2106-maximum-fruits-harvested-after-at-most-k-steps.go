package solutions

func MaxTotalFruits(fruits [][]int, startPos int, k int) int {
	n := len(fruits)
	sum := make([]int, n+1)
	indices := make([]int, n)

	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + fruits[i][1]
		indices[i] = fruits[i][0]
	}

	ans := 0
	lowerBound := func(arr []int, target int) int {
		left, right := 0, len(arr)
		for left < right {
			mid := left + (right-left)/2
			if arr[mid] < target {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return left
	}

	upperBound := func(arr []int, target int) int {
		left, right := 0, len(arr)
		for left < right {
			mid := left + (right-left)/2
			if arr[mid] <= target {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return left
	}

	for x := 0; x <= k/2; x++ {
		y := k - 2*x
		left := startPos - x
		right := startPos + y

		start := lowerBound(indices, left)
		end := upperBound(indices, right)
		ans = max(ans, sum[end]-sum[start])

		y = k - 2*x
		left = startPos - y
		right = startPos + x

		start = lowerBound(indices, left)
		end = upperBound(indices, right)
		ans = max(ans, sum[end]-sum[start])
	}

	return ans
}
