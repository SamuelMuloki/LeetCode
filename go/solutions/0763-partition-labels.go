package solutions

func PartitionLabels(s string) []int {
	last := [26]int{}
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	j, anchor := 0, 0
	ans := make([]int, 0)
	for i := 0; i < len(s); i++ {
		j = max(j, last[s[i]-'a'])
		if i == j {
			ans = append(ans, i-anchor+1)
			anchor = i + 1
		}
	}

	return ans
}
