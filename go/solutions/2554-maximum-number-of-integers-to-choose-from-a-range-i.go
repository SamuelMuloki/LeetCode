package solutions

func MaxCount(banned []int, n int, maxSum int) int {
	arr := make([]bool, n+1)
	for _, num := range banned {
		if num >= n+1 {
			continue
		}
		arr[num] = true
	}

	ans := 0
	for num := 1; num <= n; num++ {
		if arr[num] {
			continue
		}

		maxSum -= num
		if maxSum < 0 {
			break
		}
		ans++
	}

	return ans
}
