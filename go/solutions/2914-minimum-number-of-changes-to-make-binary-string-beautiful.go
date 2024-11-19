package solutions

func MinChanges(s string) int {
	ans := 0
	for i := 1; i < len(s); i += 2 {
		if s[i] != s[i-1] {
			ans++
		}
	}
	return ans
}
