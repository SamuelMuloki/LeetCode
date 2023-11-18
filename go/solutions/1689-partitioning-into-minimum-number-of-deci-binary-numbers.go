package solutions

func MinPartitions(n string) int {
	ans := 0
	for i := range n {
		ans = max(ans, int(n[i]-'0'))
	}
	return ans
}
