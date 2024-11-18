package solutions

func Decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)
	if k == 0 {
		return ans
	}

	for i := 0; i < n; i++ {
		count := 0
		for j, m := abs(k), i; j > 0; j-- {
			if k > 0 {
				count += code[(m+1)%n]
				m++
			} else {
				count += code[(n+m-1)%n]
				m--
			}
		}
		ans[i] = count
	}

	return ans
}
