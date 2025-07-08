package solutions

func LongestSubsequence(s string, k int) int {
	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			res++
		}
	}

	val := 0
	pow := 1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' {
			if val+pow > k {
				continue
			}
			val += pow
			res++
		}
		if pow > k {
			break
		}
		pow <<= 1
	}

	return res
}
