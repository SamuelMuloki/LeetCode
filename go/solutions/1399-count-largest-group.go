package solutions

func CountLargestGroup(n int) int {
	groups := make(map[int]int)
	maxCnt := 0
	for i := 1; i <= n; i++ {
		num := i
		cnt := 0
		for num > 0 {
			cnt += num % 10
			num /= 10
		}
		groups[cnt]++
		maxCnt = max(maxCnt, groups[cnt])
	}

	res := 0
	for _, num := range groups {
		if num == maxCnt {
			res++
		}
	}

	return res
}
