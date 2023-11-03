package solutions

func BuildArray(target []int, n int) []string {
	ans := make([]string, 0)
	j := 0
	for i := 1; j < len(target) && i <= n; i++ {
		if target[j] == i {
			j++
			ans = append(ans, "Push")
		} else {
			ans = append(ans, "Push")
			ans = append(ans, "Pop")
		}
	}

	return ans
}
