package solutions

func MaxCount(banned []int, n int, maxSum int) int {
	arr := make([]bool, n+1)
	for _, num := range banned {
		if num >= n+1 {
			continue
		}
		arr[num] = true
	}

	ans, start := 0, 1
	count, count2 := 0, 0
	for num := 1; num <= n; num++ {
		if arr[num] {
			continue
		}
		count, count2 = count+num, count2+1
		if count > maxSum {
			for arr[start] {
				start++
			}
			count, count2 = count-start, count2-1
		}

		ans = max(ans, count2)
	}

	return ans
}
