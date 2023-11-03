package solutions

func BuildArray(target []int, n int) []string {
	ans := make([]string, 0)
	j := 0
	for i := range target {
		for j < target[i]-1 {
			ans = append(ans, "Push")
			ans = append(ans, "Pop")
			j++
		}

		ans = append(ans, "Push")
		j++
	}

	return ans
}
