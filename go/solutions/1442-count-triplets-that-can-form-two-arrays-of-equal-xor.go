package solutions

func CountTriplets(arr []int) int {
	n := len(arr)
	ans := 0
	for i := 0; i < n; i++ {
		k := arr[i]
		for j := i + 1; j < n; j++ {
			k = k ^ arr[j]
			if k == 0 {
				ans += j - i
			}
		}
	}

	return ans
}
