package solutions

func FindClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-k

	for l < r {
		mid := l + (r-l)/2
		if x-arr[mid] > arr[mid+k]-x {
			l = mid + 1
		} else {
			r = mid
		}
	}

	closest := make([]int, k)
	for i, j := l, 0; i < l+k; i, j = i+1, j+1 {
		closest[j] = arr[i]
	}

	return closest
}
