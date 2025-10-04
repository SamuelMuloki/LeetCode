package solutions

func ReverseStr(s string, k int) string {
	res := []rune(s)
	n := len(s) - 1
	for i := 0; i < len(s)-1; i += 2 * k {
		for j, k := i, min(n, (i+k)-1); j < k; j, k = j+1, k-1 {
			res[j], res[k] = res[k], res[j]
		}
	}
	return string(res)
}
