package solutions

func SumZero(n int) []int {
	res := make([]int, n)
	num := -n / 2
	for i := 0; i <= n/2; i++ {
		if num == 0 && n%2 == 0 {
			continue
		}
		res[i] = num
		res[n-i-1] = -num
		num++
	}

	return res
}
