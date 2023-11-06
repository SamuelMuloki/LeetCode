package solutions

func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	st, ans := []int{}, make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && temperatures[st[len(st)-1]] <= temperatures[i] {
			st = st[:len(st)-1]
		}

		if len(st) > 0 {
			ans[i] = st[len(st)-1] - i
		}

		st = append(st, i)
	}

	return ans
}
